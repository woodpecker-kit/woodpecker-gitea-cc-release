package gitea_cc_plugin

import (
	"code.gitea.io/sdk/gitea"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"net"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"sync"
	"time"
)

var (
	ErrMissingTag              = fmt.Errorf("NewReleaseClientByWoodpecker missing tag, please check now in tag build")
	ErrPackageNotExist         = fmt.Errorf("PackageFetch not exist, code 404")
	ErrPathCanNotLoadGoModFile = fmt.Errorf("path can not load go.mod")
	ErrPackageGoExists         = fmt.Errorf("package go exists")
)

// Release holds ties the drone env data and gitea client together.
type releaseClient struct {
	client *gitea.Client
	debug  bool
	dryRun bool

	url        string
	ctx        context.Context
	mutex      *sync.RWMutex
	httpClient *http.Client

	accessToken string // this not in RWLock
	username    string
	password    string
	otp         string
	sudo        string

	owner string
	repo  string
	tag   string
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

	SetOTP(otp string)

	SetSudo(sudo string)

	SetBasicAuth(username, password string)

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
	if info.Build.Event != wd_info.EventPipelineTag {
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

	httpClient := &http.Client{
		Timeout: time.Duration(config.GiteaTimeoutSecond) * time.Second,
	}
	if config.GiteaInsecure {
		cookieJar, _ := cookiejar.New(nil)
		httpClient = &http.Client{
			Jar: cookieJar,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				DialContext: (&net.Dialer{
					Timeout:   time.Duration(config.TimeoutSecond*3) * time.Second,
					KeepAlive: time.Duration(config.TimeoutSecond*3) * time.Second,
				}).DialContext,
				TLSHandshakeTimeout:   time.Duration(config.TimeoutSecond) * time.Second,
				ResponseHeaderTimeout: time.Duration(config.TimeoutSecond) * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
			Timeout: time.Duration(config.GiteaTimeoutSecond) * time.Second,
		}
	}

	client, errNewClient := gitea.NewClient(config.GiteaBaseUrl,
		gitea.SetToken(config.GiteaApiKey),
		gitea.SetHTTPClient(httpClient),
	)
	if errNewClient != nil {
		return nil, fmt.Errorf("failed to create gitea client: %s", errNewClient)
	}
	wd_log.Debug("gitea client created success")

	// if the title was not provided via we use the tag instead
	if config.GiteaReleaseTitle == "" {
		config.GiteaReleaseTitle = info.Build.Tag
	}

	return &releaseClient{
		client: client,
		debug:  config.Debug,
		dryRun: config.DryRun,

		url:         strings.TrimSuffix(config.GiteaBaseUrl, "/"),
		ctx:         context.Background(),
		mutex:       &sync.RWMutex{},
		httpClient:  httpClient,
		accessToken: config.GiteaApiKey,

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
	}, nil
}
