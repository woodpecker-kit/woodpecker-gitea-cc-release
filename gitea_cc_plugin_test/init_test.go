package gitea_cc_plugin_test

import (
	"fmt"
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/sinlov-go/unittest-kit/unittest_file_kit"
	"github.com/woodpecker-kit/woodpecker-gitea-cc-release/gitea_cc_plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_steps_transfer"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	keyEnvDebug  = "CI_DEBUG"
	keyEnvCiNum  = "CI_NUMBER"
	keyEnvCiKey  = "CI_KEY"
	keyEnvCiKeys = "CI_KEYS"

	mockVersion = "v1.0.0"
	mockName    = "woodpecker-gitea-cc-release"
)

var (
	// testBaseFolderPath
	//  test base dir will auto get by package init()
	testBaseFolderPath = ""
	testGoldenKit      *unittest_file_kit.TestGoldenKit

	// mustSetInCiEnvList
	//  for check set in CI env not empty
	mustSetInCiEnvList = []string{
		wd_flag.EnvKeyCiSystemPlatform,
		wd_flag.EnvKeyCiSystemVersion,
	}
	// mustSetArgsAsEnvList
	mustSetArgsAsEnvList = []string{
		gitea_cc_plugin.EnvGiteaBaseUrl,
		gitea_cc_plugin.EnvGiteaApiKey,
	}

	valEnvTimeoutSecond          uint
	valEnvPluginDebug            = false
	valEnvGiteaDryRun            = true
	valEnvGiteaDraft             = false
	valEnvGiteaPrerelease        = true
	valEnvGiteaBaseUrl           = ""
	valEnvGiteaInsecure          = false
	valEnvGiteaApiKey            = ""
	valEnvGiteaReleaseFilesGlobs []string
	valEnvGiteaFileExistsDo      = gitea_cc_plugin.EnvGiteaDraft
	valEnvGiteaFilesChecksum     []string
)

func init() {
	testBaseFolderPath, _ = getCurrentFolderPath()
	wd_log.SetLogLineDeep(2)
	// if open wd_template please open this
	//wd_template.RegisterSettings(wd_template.DefaultHelpers)

	testGoldenKit = unittest_file_kit.NewTestGoldenKit(testBaseFolderPath)

	valEnvTimeoutSecond = uint(env_kit.FetchOsEnvInt(wd_flag.EnvKeyPluginTimeoutSecond, 10))
	valEnvPluginDebug = env_kit.FetchOsEnvBool(wd_flag.EnvKeyPluginDebug, false)
	valEnvGiteaDryRun = env_kit.FetchOsEnvBool(gitea_cc_plugin.EnvGiteaDryRun, true)
	valEnvGiteaDraft = env_kit.FetchOsEnvBool(gitea_cc_plugin.EnvGiteaDraft, false)
	valEnvGiteaPrerelease = env_kit.FetchOsEnvBool(gitea_cc_plugin.EnvGiteaPrerelease, true)
	valEnvGiteaBaseUrl = env_kit.FetchOsEnvStr(gitea_cc_plugin.EnvGiteaBaseUrl, "")
	valEnvGiteaInsecure = env_kit.FetchOsEnvBool(gitea_cc_plugin.EnvGiteaInsecure, false)
	valEnvGiteaApiKey = env_kit.FetchOsEnvStr(gitea_cc_plugin.EnvGiteaApiKey, "")
	valEnvGiteaReleaseFilesGlobs = env_kit.FetchOsEnvStringSlice(gitea_cc_plugin.EnvGiteaReleaseFilesGlobs)
	valEnvGiteaFileExistsDo = env_kit.FetchOsEnvStr(gitea_cc_plugin.EnvGiteaReleaseFileExistsDo, gitea_cc_plugin.FileExistsDoFail)
	valEnvGiteaFilesChecksum = env_kit.FetchOsEnvStringSlice(gitea_cc_plugin.EnvGiteaFilesChecksum)
}

// test case basic tools start
// getCurrentFolderPath
//
//	can get run path this golang dir
func getCurrentFolderPath() (string, error) {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "", fmt.Errorf("can not get current file info")
	}
	return filepath.Dir(file), nil
}

// test case basic tools end

func envCheck(t *testing.T) bool {

	if valEnvPluginDebug {
		wd_log.OpenDebug()
	}

	// most CI system will set env CI to true
	envCI := env_kit.FetchOsEnvStr("CI", "")
	if envCI == "" {
		t.Logf("not in CI system, skip envCheck")
		return false
	}
	t.Logf("check env for CI system")
	return env_kit.MustHasEnvSetByArray(t, mustSetInCiEnvList)
}

func envMustArgsCheck(t *testing.T) bool {
	for _, item := range mustSetArgsAsEnvList {
		if os.Getenv(item) == "" {
			t.Logf("plasee set env: %s, than run test\nfull need set env %v", item, mustSetArgsAsEnvList)
			return true
		}
	}
	return false
}

func generateTransferStepsOut(plugin gitea_cc_plugin.Plugin, mark string, data interface{}) error {
	_, err := wd_steps_transfer.Out(plugin.Settings.RootPath, plugin.Settings.StepsTransferPath, plugin.GetWoodPeckerInfo(), mark, data)
	return err
}

func mockPluginSettings() gitea_cc_plugin.Settings {
	// all mock settings can set here
	settings := gitea_cc_plugin.Settings{
		// use env:PLUGIN_DEBUG
		Debug:             valEnvPluginDebug,
		TimeoutSecond:     valEnvTimeoutSecond,
		RootPath:          testGoldenKit.GetTestDataFolderFullPath(),
		StepsTransferPath: wd_steps_transfer.DefaultKitStepsFileName,
	}
	settings.DryRun = valEnvGiteaDryRun
	settings.GiteaDraft = valEnvGiteaDraft
	settings.GiteaPrerelease = valEnvGiteaPrerelease
	settings.GiteaBaseUrl = valEnvGiteaBaseUrl
	settings.GiteaInsecure = valEnvGiteaInsecure
	settings.GiteaApiKey = valEnvGiteaApiKey
	settings.GiteaReleaseFilesGlobs = valEnvGiteaReleaseFilesGlobs
	settings.GiteaReleaseFileExistsDo = valEnvGiteaFileExistsDo
	if settings.GiteaReleaseFileExistsDo == "" {
		settings.GiteaReleaseFileExistsDo = gitea_cc_plugin.FileExistsDoFail
	}
	settings.GiteaFilesChecksum = valEnvGiteaFilesChecksum

	return settings

}

func mockPluginWithSettings(t *testing.T, woodpeckerInfo wd_info.WoodpeckerInfo, settings gitea_cc_plugin.Settings) gitea_cc_plugin.Plugin {
	p := gitea_cc_plugin.Plugin{
		Name:    mockName,
		Version: mockVersion,
	}

	// mock woodpecker info
	//t.Log("mockPluginWithStatus")
	p.SetWoodpeckerInfo(woodpeckerInfo)

	if p.ShortInfo().Build.WorkSpace != "" {
		settings.RootPath = p.ShortInfo().Build.WorkSpace
	}

	p.Settings = settings
	return p
}
