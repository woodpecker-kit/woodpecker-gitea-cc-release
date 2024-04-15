package gitea_cc_plugin

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
)

const (
	CliNameGiteaDryRun = "settings.gitea-dry-run"
	EnvGiteaDryRun     = "PLUGIN_GITEA_DRY_RUN"

	CliNameGiteaDraft = "settings.gitea-draft"
	EnvGiteaDraft     = "PLUGIN_GITEA_DRAFT"

	CliNameGiteaPrerelease = "settings.gitea-prerelease"
	EnvGiteaPrerelease     = "PLUGIN_GITEA_PRERELEASE"

	CliNameGiteaBaseUrl = "settings.gitea-base-url"
	EnvGiteaBaseUrl     = "PLUGIN_GITEA_BASE_URL"

	CliNameGiteaInsecure = "settings.gitea-insecure"
	EnvGiteaInsecure     = "PLUGIN_GITEA_INSECURE"

	CliNameGiteaApiKey = "settings.gitea-api-key"
	EnvGiteaApiKey     = "PLUGIN_GITEA_API_KEY"

	CliNameGiteaReleaseFilesGlobs = "settings.gitea-release-files-globs"
	EnvGiteaReleaseFilesGlobs     = "PLUGIN_GITEA_RELEASE_FILES_GLOBS"

	CliNameGiteaReleaseFileExistsDo = "settings.gitea-release-file-exists-do"
	EnvGiteaReleaseFileExistsDo     = "PLUGIN_GITEA_RELEASE_FILE_EXISTS_DO"

	CliNameGiteaFilesChecksum = "settings.gitea-files-checksum"
	EnvGiteaFilesChecksum     = "PLUGIN_GITEA_FILES_CHECKSUM"

	// remove or change this code

	CliNameNotEmptyEnvs = "settings.not-empty-envs"
	EnvNotEmptyEnvs     = "PLUGIN_NOT_EMPTY_ENVS"

	CliNamePrinterPrintKeys = "settings.env-printer-print-keys"
	EnvPrinterPrintKeys     = "PLUGIN_ENV_PRINTER_PRINT_KEYS"

	CliNamePrinterPaddingLeftMax = "settings.env-printer-padding-left-max"
	EnvPrinterPaddingLeftMax     = "PLUGIN_ENV_PRINTER_PADDING_LEFT_MAX"

	CliNameStepsTransferDemo = "settings.steps-transfer-demo"
	EnvStepsTransferDemo     = "PLUGIN_STEPS_TRANSFER_DEMO"
)

// GlobalFlag
// Other modules also have flags
func GlobalFlag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    CliNameGiteaDryRun,
			Usage:   "gitea release dry run",
			Value:   true,
			EnvVars: []string{EnvGiteaDryRun},
		},
		&cli.BoolFlag{
			Name:    CliNameGiteaDraft,
			Usage:   "gitea release draft",
			Value:   false,
			EnvVars: []string{EnvGiteaDraft},
		},
		&cli.BoolFlag{
			Name:    CliNameGiteaPrerelease,
			Usage:   "gitea release type prerelease",
			Value:   true,
			EnvVars: []string{EnvGiteaPrerelease},
		},
		&cli.StringFlag{
			Name:    CliNameGiteaBaseUrl,
			Usage:   "gitea base url, Required",
			EnvVars: []string{EnvGiteaBaseUrl},
		},
		&cli.BoolFlag{
			Name:    CliNameGiteaInsecure,
			Usage:   "visit base-url via insecure https protocol",
			Value:   false,
			EnvVars: []string{EnvGiteaInsecure},
		},
		&cli.StringFlag{
			Name:    CliNameGiteaApiKey,
			Usage:   "gitea api key, Required",
			EnvVars: []string{EnvGiteaApiKey},
		},

		// release files
		&cli.StringSliceFlag{
			Name:    CliNameGiteaReleaseFilesGlobs,
			Usage:   "release as files by glob pattern, if empty will skip release files",
			EnvVars: []string{EnvGiteaReleaseFilesGlobs},
		},
		&cli.StringFlag{
			Name:    CliNameGiteaReleaseFileExistsDo,
			Usage:   fmt.Sprintf("what to do if release file already exist support: %v", supportFileExistsDoList),
			Value:   FileExistsDoFail,
			EnvVars: []string{EnvGiteaReleaseFileExistsDo},
		},
		&cli.StringSliceFlag{
			Name:    CliNameGiteaFilesChecksum,
			Usage:   fmt.Sprintf("generate specific checksums, empty will skip, support: %v", CheckSumSupport),
			EnvVars: []string{EnvGiteaFilesChecksum},
		},
	}
}

func HideGlobalFlag() []cli.Flag {
	return []cli.Flag{}
}

func BindCliFlags(c *cli.Context,
	debug bool,
	cliName, cliVersion string,
	wdInfo *wd_info.WoodpeckerInfo,
	rootPath,
	stepsTransferPath string, stepsOutDisable bool,
) (*Plugin, error) {

	config := Settings{
		Debug:             debug,
		TimeoutSecond:     c.Uint(wd_flag.NameCliPluginTimeoutSecond),
		StepsTransferPath: stepsTransferPath,
		StepsOutDisable:   stepsOutDisable,
		RootPath:          rootPath,

		DryRun:          c.Bool(CliNameGiteaDryRun),
		GiteaDraft:      c.Bool(CliNameGiteaDraft),
		GiteaPrerelease: c.Bool(CliNameGiteaPrerelease),
		GiteaBaseUrl:    c.String(CliNameGiteaBaseUrl),
		GiteaInsecure:   c.Bool(CliNameGiteaInsecure),
		GiteaApiKey:     c.String(CliNameGiteaApiKey),

		GiteaReleaseFilesGlobs:   c.StringSlice(CliNameGiteaReleaseFilesGlobs),
		GiteaReleaseFileExistsDo: c.String(CliNameGiteaReleaseFileExistsDo),
		GiteaFilesChecksum:       c.StringSlice(CliNameGiteaFilesChecksum),
	}

	// set default TimeoutSecond
	if config.TimeoutSecond == 0 {
		config.TimeoutSecond = 10
	}

	wd_log.Debugf("args %s: %v", wd_flag.NameCliPluginTimeoutSecond, config.TimeoutSecond)

	infoShort := wd_short_info.ParseWoodpeckerInfo2Short(*wdInfo)

	p := Plugin{
		Name:           cliName,
		Version:        cliVersion,
		woodpeckerInfo: wdInfo,
		wdShortInfo:    &infoShort,
		Settings:       config,
	}

	return &p, nil
}
