package gitea_cc_plugin

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"path/filepath"
)

const (
	CliNameGiteaApiKey = "settings.gitea-api-key"
	EnvGiteaApiKey     = "PLUGIN_GITEA_API_KEY"

	CliNameGiteaBaseUrl = "settings.gitea-base-url"
	EnvGiteaBaseUrl     = "PLUGIN_GITEA_BASE_URL"

	CliNameGiteaInsecure = "settings.gitea-insecure"
	EnvGiteaInsecure     = "PLUGIN_GITEA_INSECURE"

	CliNameGiteaDryRun = "settings.gitea-dry-run"
	EnvGiteaDryRun     = "PLUGIN_GITEA_DRY_RUN"

	CliNameGiteaDraft = "settings.gitea-draft"
	EnvGiteaDraft     = "PLUGIN_GITEA_DRAFT"

	CliNameGiteaPrerelease = "settings.gitea-prerelease"
	EnvGiteaPrerelease     = "PLUGIN_GITEA_PRERELEASE"

	CliNameGiteaReleaseFileRootPath = "settings.gitea-release-file-root-path"
	EnvGiteaReleaseFileRootPath     = "PLUGIN_GITEA_RELEASE_FILE_ROOT_PATH"

	CliNameGiteaReleaseFilesGlobs = "settings.gitea-release-files-globs"
	EnvGiteaReleaseFilesGlobs     = "PLUGIN_GITEA_RELEASE_FILES_GLOBS"

	CliNameGiteaReleaseFileExistsDo = "settings.gitea-release-file-exists-do"
	EnvGiteaReleaseFileExistsDo     = "PLUGIN_GITEA_RELEASE_FILE_EXISTS_DO"

	CliNameGiteaFilesChecksum = "settings.gitea-files-checksum"
	EnvGiteaFilesChecksum     = "PLUGIN_GITEA_FILES_CHECKSUM"

	CliNameGiteaReleaseTitle = "settings.gitea-release-title"
	EnvGiteaReleaseTitle     = "PLUGIN_GITEA_RELEASE_TITLE"

	CliNameGiteaReleaseNote = "settings.gitea-release-note"
	EnvGiteaReleaseNote     = "PLUGIN_GITEA_RELEASE_NOTE"

	CliNameGiteaReleaseNoteByConventionChange = "settings.gitea-release-note-by-convention-change"
	EnvGiteaReleaseNoteByConventionChange     = "PLUGIN_GITEA_RELEASE_NOTE_BY_CONVENTION_CHANGE"

	CliNameGiteaReleaseReadChangeLogFile = "settings.gitea-release-read-change-log-file"
	EnvGiteaReleaseReadChangeLogFile     = "PLUGIN_GITEA_RELEASE_READ_CHANGE_LOG_FILE"

	CliNameGiteaTimeoutSecond = "settings.gitea-timeout-second"
	EnvGiteaTimeoutSecond     = "PLUGIN_GITEA_TIMEOUT_SECOND"
)

// GlobalFlag
// Other modules also have flags
func GlobalFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    CliNameGiteaApiKey,
			Usage:   "gitea api key, Required",
			EnvVars: []string{EnvGiteaApiKey},
		},
		&cli.StringFlag{
			Name:    CliNameGiteaBaseUrl,
			Usage:   fmt.Sprintf("gitea base url, when `%s` is `gitea`, and this flag is empty, will try get from `%s`", wd_flag.EnvKeyCiForgeType, wd_flag.EnvKeyCiForgeUrl),
			EnvVars: []string{EnvGiteaBaseUrl},
		},
		&cli.BoolFlag{
			Name:    CliNameGiteaInsecure,
			Usage:   "visit base-url via insecure https protocol",
			Value:   false,
			EnvVars: []string{EnvGiteaInsecure},
		},

		&cli.BoolFlag{
			Name:    CliNameGiteaDryRun,
			Usage:   "gitea release dry run",
			Value:   false,
			EnvVars: []string{EnvGiteaDryRun},
		},
		&cli.BoolFlag{
			Name:    CliNameGiteaDraft,
			Usage:   "gitea release setting: type draft",
			Value:   false,
			EnvVars: []string{EnvGiteaDraft},
		},
		&cli.BoolFlag{
			Name:    CliNameGiteaPrerelease,
			Usage:   "gitea release setting: type prerelease",
			Value:   true,
			EnvVars: []string{EnvGiteaPrerelease},
		},

		// release files
		&cli.StringFlag{
			Name:    CliNameGiteaReleaseFileRootPath,
			Usage:   "release file root path, if empty will use workspace, most of use project root path",
			EnvVars: []string{EnvGiteaReleaseFileRootPath},
		},
		&cli.StringSliceFlag{
			Name:    CliNameGiteaReleaseFilesGlobs,
			Usage:   "release as files by glob pattern, if empty will skip release files",
			EnvVars: []string{EnvGiteaReleaseFilesGlobs},
		},
		&cli.StringFlag{
			Name:    CliNameGiteaReleaseFileExistsDo,
			Usage:   fmt.Sprintf("do if release update file already exist, support: %v", supportFileExistsDoList),
			Value:   FileExistsDoFail,
			EnvVars: []string{EnvGiteaReleaseFileExistsDo},
		},
		&cli.StringSliceFlag{
			Name:    CliNameGiteaFilesChecksum,
			Usage:   fmt.Sprintf("generate specific checksums, empty will skip, support: %v", CheckSumSupport),
			EnvVars: []string{EnvGiteaFilesChecksum},
		},

		// release info
		&cli.StringFlag{
			Name:    CliNameGiteaReleaseTitle,
			Usage:   "release title, if empty will use tag, can be cover by tag name of convention change log",
			EnvVars: []string{EnvGiteaReleaseTitle},
		},
		&cli.StringFlag{
			Name:    CliNameGiteaReleaseNote,
			Usage:   "release note, can be cover by tag name of convention change log",
			EnvVars: []string{EnvGiteaReleaseNote},
		},
		&cli.BoolFlag{
			Name:    CliNameGiteaReleaseNoteByConventionChange,
			Usage:   fmt.Sprintf("release note by convention change, if true will read change log file, and cover flag %s", CliNameGiteaReleaseNote),
			EnvVars: []string{EnvGiteaReleaseNoteByConventionChange},
		},
		&cli.StringFlag{
			Name:    CliNameGiteaReleaseReadChangeLogFile,
			Usage:   fmt.Sprintf("release read change log file, if empty will use default %s", versionConventionChangeLogFileName),
			Value:   versionConventionChangeLogFileName,
			EnvVars: []string{EnvGiteaReleaseReadChangeLogFile},
		},
	}
}

func HideGlobalFlag() []cli.Flag {
	return []cli.Flag{
		&cli.UintFlag{
			Name:    CliNameGiteaTimeoutSecond,
			Usage:   "gitea release api timeout second, default 60, less 30",
			Value:   60,
			Hidden:  true,
			EnvVars: []string{EnvGiteaTimeoutSecond},
		},
	}
}

func BindCliFlags(c *cli.Context,
	debug bool,
	cliName, cliVersion string,
	wdInfo *wd_info.WoodpeckerInfo,
	rootPath,
	stepsTransferPath string, stepsOutDisable bool,
) (*GiteaCCRelease, error) {

	releaseFileRootPath := c.String(CliNameGiteaReleaseFileRootPath)
	if releaseFileRootPath == "" {
		releaseFileRootPath = rootPath
	}

	readCCLogFileName := c.String(CliNameGiteaReleaseReadChangeLogFile)
	if readCCLogFileName == "" { // empty will use default
		readCCLogFileName = versionConventionChangeLogFileName
	}
	changeLogFullPath := filepath.Join(rootPath, readCCLogFileName)

	config := Settings{
		Debug:             debug,
		TimeoutSecond:     c.Uint(wd_flag.NameCliPluginTimeoutSecond),
		StepsTransferPath: stepsTransferPath,
		StepsOutDisable:   stepsOutDisable,
		RootPath:          rootPath,

		GiteaBaseUrl:       c.String(CliNameGiteaBaseUrl),
		GiteaInsecure:      c.Bool(CliNameGiteaInsecure),
		GiteaApiKey:        c.String(CliNameGiteaApiKey),
		GiteaTimeoutSecond: c.Uint(CliNameGiteaTimeoutSecond),

		GiteaPrerelease: c.Bool(CliNameGiteaPrerelease),
		DryRun:          c.Bool(CliNameGiteaDryRun),
		GiteaDraft:      c.Bool(CliNameGiteaDraft),

		GiteaReleaseFilesGlobs:       c.StringSlice(CliNameGiteaReleaseFilesGlobs),
		GiteaReleaseFileGlobRootPath: releaseFileRootPath,
		GiteaReleaseFileExistsDo:     c.String(CliNameGiteaReleaseFileExistsDo),
		GiteaFilesChecksum:           c.StringSlice(CliNameGiteaFilesChecksum),

		GiteaReleaseTitle:                  c.String(CliNameGiteaReleaseTitle),
		GiteaReleaseNote:                   c.String(CliNameGiteaReleaseNote),
		GiteaReleaseNoteByConventionChange: c.Bool(CliNameGiteaReleaseNoteByConventionChange),
		GiteaReleaseConventionReadPath:     changeLogFullPath,
	}

	// set default TimeoutSecond
	if config.TimeoutSecond == 0 {
		config.TimeoutSecond = 10
	}

	if config.GiteaTimeoutSecond < 30 {
		config.GiteaTimeoutSecond = 30
	}

	wd_log.Debugf("args %s: %v", wd_flag.NameCliPluginTimeoutSecond, config.TimeoutSecond)

	infoShort := wd_short_info.ParseWoodpeckerInfo2Short(*wdInfo)

	p := GiteaCCRelease{
		Name:           cliName,
		Version:        cliVersion,
		woodpeckerInfo: wdInfo,
		wdShortInfo:    &infoShort,
		Settings:       config,
	}

	return &p, nil
}
