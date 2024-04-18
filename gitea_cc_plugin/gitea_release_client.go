package gitea_cc_plugin

import (
	"code.gitea.io/sdk/gitea"
	"fmt"
	"github.com/sinlov-go/gitea-client-wrapper/gitea_token_client"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"path/filepath"
	"strings"
)

var (
	ErrMissingTag = fmt.Errorf("NewReleaseClientByWoodpecker missing tag, please check now in tag build")
)

// Release holds ties the drone env data and gitea client together.
type releaseClient struct {
	gitea_token_client.GiteaTokenClient

	dryRun bool
	owner  string
	repo   string
	tag    string
	// tagTarget
	//is the branch or commit sha to tag
	tagTarget  string
	draft      bool
	prerelease bool
	// what to do if file already exist can use: overwrite, skip
	fileExistsDo string
	title        string
	note         string

	uploadFilePaths []string
	uploadDesc      string
}

type GiteaPackageInfo struct {
	Id      uint64 `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type PluginGiteaReleaseClient interface {
	GetUploadDesc() string

	Title() string

	SetTitle(title string)

	Tag() string

	SetNote(noteContent string)

	BuildRelease() (*gitea.Release, error)

	UploadFiles(releaseID int64) error
}

// NewReleaseClientByWoodpeckerShort creates a new release client from the woodpecker info.
// will try to upload files if the globs are set.
// and will create a release with the tag as the title.
// check sums can be generated for the files [ sum.txt ] from flag CliNameGiteaReleaseFileRootPath
func NewReleaseClientByWoodpeckerShort(info wd_short_info.WoodpeckerInfoShort, config Settings) (PluginGiteaReleaseClient, error) {
	if info.Build.Event != wd_info.EventPipelineTag && !config.DryRun {
		return nil, ErrMissingTag
	}
	uploadDesc := ""
	var uploadFiles []string
	if len(config.GiteaReleaseFilesGlobs) > 0 {
		findFiles, errGlobs := FindFileByGlobs(config.GiteaReleaseFilesGlobs, config.GiteaReleaseFileGlobRootPath)
		if errGlobs != nil {
			return nil, errGlobs
		}

		if len(findFiles) == 0 {
			return nil, fmt.Errorf("not found files by globs: %v , at path: %s", config.GiteaReleaseFilesGlobs, config.GiteaReleaseFileGlobRootPath)
		}

		repetitionFiles := findUploadFileRepetitionByBaseName(findFiles)
		if len(repetitionFiles) > 0 {
			return nil, fmt.Errorf("found files repetition by base name, now not support upload, repetition path as\n%s", strings.Join(repetitionFiles, "\n"))
		}

		uploadFiles = findFiles

		if len(config.GiteaFilesChecksum) > 0 {
			filesCheckRes, errCheckSum := WriteChecksumsByFiles(uploadFiles, config.GiteaFilesChecksum, config.GiteaReleaseFileGlobRootPath)

			if errCheckSum != nil {
				errCheckSumWrite := fmt.Errorf("from config.files_checksum failed: %v", errCheckSum)
				return nil, errCheckSumWrite
			}
			uploadFiles = filesCheckRes
		}
	} else {
		uploadDesc = "PLUGIN_RELEASE_GITEA_FILES not setting, skip upload files"
	}

	// if the title was not provided via we use the tag instead
	if config.GiteaReleaseTitle == "" {
		if info.Build.Event == wd_info.EventPipelineTag {
			config.GiteaReleaseTitle = info.Build.Tag
		} else {
			config.GiteaReleaseTitle = defaultTitle
		}
	}

	rc := &releaseClient{
		dryRun:       config.DryRun,
		owner:        info.Repo.OwnerName,
		repo:         info.Repo.ShortName,
		tag:          info.Build.Tag,
		tagTarget:    info.Build.CommitBranch,
		draft:        config.GiteaDraft,
		prerelease:   config.GiteaPrerelease,
		fileExistsDo: config.GiteaReleaseFileExistsDo,
		title:        config.GiteaReleaseTitle,
		note:         config.GiteaReleaseNote,

		uploadFilePaths: uploadFiles,
		uploadDesc:      uploadDesc,
	}
	errNewClient := rc.NewClientWithHttpTimeout(config.GiteaBaseUrl, config.GiteaApiKey, config.GiteaTimeoutSecond, config.GiteaInsecure)
	if errNewClient != nil {
		return nil, errNewClient
	}
	wd_log.Debug("gitea client created success")
	return rc, nil
}

// findUploadFileRepetitionByBaseName find upload file base name repetition
// return duplicates file full path, len 0 is not find
func findUploadFileRepetitionByBaseName(files []string) []string {
	seen := make(map[string]bool)
	var duplicates []string
	for _, fileFullPath := range files {
		fileBaseName := filepath.Base(fileFullPath)
		if seen[fileBaseName] {
			duplicates = append(duplicates, fileFullPath)
		} else {
			seen[fileBaseName] = true
		}
	}
	return duplicates
}
