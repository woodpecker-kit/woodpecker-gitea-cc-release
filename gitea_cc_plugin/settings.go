package gitea_cc_plugin

const (
	FileExistsDoFail      = "fail"
	FileExistsDoOverwrite = "overwrite"
	FileExistsDoSkip      = "skip"

	defaultTitle = "Release title"

	// versionRcFileName read convention change log config
	versionRcFileName = ".versionrc"

	// versionConventionChangeLogFileName for read convention change log file name
	versionConventionChangeLogFileName = "CHANGELOG.md"
)

type (
	// Settings gitea_cc_plugin private config
	Settings struct {
		Debug             bool
		TimeoutSecond     uint
		StepsTransferPath string
		StepsOutDisable   bool
		RootPath          string

		DryRun             bool
		GiteaDraft         bool
		GiteaPrerelease    bool
		GiteaTimeoutSecond uint
		GiteaBaseUrl       string
		GiteaInsecure      bool
		GiteaApiKey        string

		GiteaReleaseFilesGlobs       []string
		GiteaReleaseFileGlobRootPath string
		GiteaReleaseFileExistsDo     string
		GiteaFilesChecksum           []string

		GiteaReleaseTitle                  string
		GiteaReleaseNote                   string
		GiteaReleaseNoteByConventionChange bool
		GiteaReleaseConventionReadPath     string
	}
)

var (
	supportFileExistsDoList = []string{
		FileExistsDoFail,
		FileExistsDoOverwrite,
		FileExistsDoSkip,
	}
)
