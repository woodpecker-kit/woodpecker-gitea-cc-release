package gitea_cc_plugin

const (
	FileExistsDoFail      = "fail"
	FileExistsDoOverwrite = "overwrite"
	FileExistsDoSkip      = "skip"

	// StepsTransferMarkDemoConfig
	// steps transfer key
	StepsTransferMarkDemoConfig = "demo_config"
)

type (
	// Settings gitea_cc_plugin private config
	Settings struct {
		Debug             bool
		TimeoutSecond     uint
		StepsTransferPath string
		StepsOutDisable   bool
		RootPath          string

		DryRun          bool
		GiteaDraft      bool
		GiteaPrerelease bool
		GiteaBaseUrl    string
		GiteaInsecure   bool
		GiteaApiKey     string

		GiteaReleaseFilesGlobs   []string
		GiteaReleaseFileExistsDo string
		GiteaFilesChecksum       []string
	}
)

var (
	supportFileExistsDoList = []string{
		FileExistsDoFail,
		FileExistsDoOverwrite,
		FileExistsDoSkip,
	}
)
