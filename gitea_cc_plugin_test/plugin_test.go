package gitea_cc_plugin_test

import (
	"github.com/woodpecker-kit/woodpecker-gitea-cc-release/gitea_cc_plugin"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"path/filepath"
	"testing"
)

func TestCheckArgsPlugin(t *testing.T) {
	t.Log("mock GiteaCCRelease")
	// successArgs
	successArgsWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	successArgsSettings := mockPluginSettings()
	successArgsSettings.GiteaBaseUrl = "foo url"
	successArgsSettings.GiteaApiKey = "bar key"

	// baseUrlEmpty
	baseUrlEmptyWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	baseUrlEmptySettings := mockPluginSettings()
	baseUrlEmptySettings.GiteaApiKey = "bar key"

	// apkKeyEmpty
	apkKeyEmptyWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	apkKeyEmptySettings := mockPluginSettings()
	apkKeyEmptySettings.GiteaBaseUrl = "foo url"

	// filesExistsDoError
	filesExistsDoErrorWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	filesExistsDoErrorSettings := mockPluginSettings()
	filesExistsDoErrorSettings.GiteaBaseUrl = "foo url"
	filesExistsDoErrorSettings.GiteaApiKey = "bar key"
	filesExistsDoErrorSettings.GiteaReleaseFilesGlobs = []string{"*.zip"}
	filesExistsDoErrorSettings.GiteaReleaseFileExistsDo = "error"

	// fileChecksumError
	fileChecksumErrorWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
	)
	fileChecksumErrorSettings := mockPluginSettings()
	fileChecksumErrorSettings.GiteaBaseUrl = "foo url"
	fileChecksumErrorSettings.GiteaApiKey = "bar key"
	fileChecksumErrorSettings.GiteaReleaseFilesGlobs = []string{"*.zip"}
	fileChecksumErrorSettings.GiteaFilesChecksum = []string{"some"}

	tests := []struct {
		name           string
		woodpeckerInfo wd_info.WoodpeckerInfo
		settings       gitea_cc_plugin.Settings
		workRoot       string

		isDryRun          bool
		wantArgFlagNotErr bool
	}{
		{
			name:              "successArgs",
			woodpeckerInfo:    successArgsWoodpeckerInfo,
			settings:          successArgsSettings,
			wantArgFlagNotErr: true,
		},
		{
			name:           "baseUrlEmpty",
			woodpeckerInfo: baseUrlEmptyWoodpeckerInfo,
			settings:       baseUrlEmptySettings,
		},
		{
			name:           "apkKeyEmpty",
			woodpeckerInfo: apkKeyEmptyWoodpeckerInfo,
			settings:       apkKeyEmptySettings,
		},
		{
			name:           "filesExistsDoError",
			woodpeckerInfo: filesExistsDoErrorWoodpeckerInfo,
			settings:       filesExistsDoErrorSettings,
		},
		{
			name:           "fileChecksumError",
			woodpeckerInfo: fileChecksumErrorWoodpeckerInfo,
			settings:       fileChecksumErrorSettings,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := mockPluginWithSettings(t, tc.woodpeckerInfo, tc.settings)
			p.OnlyArgsCheck()
			errPluginRun := p.Exec()
			if tc.wantArgFlagNotErr {
				if errPluginRun != nil {
					wdShotInfo := wd_short_info.ParseWoodpeckerInfo2Short(p.GetWoodPeckerInfo())
					wd_log.VerboseJsonf(wdShotInfo, "print WoodpeckerInfoShort")
					wd_log.VerboseJsonf(p.Settings, "print Settings")
					t.Fatalf("wantArgFlagNotErr %v\np.Exec() error:\n%v", tc.wantArgFlagNotErr, errPluginRun)
					return
				}
				infoShot := p.ShortInfo()
				wd_log.VerboseJsonf(infoShot, "print WoodpeckerInfoShort")
			} else {
				if errPluginRun == nil {
					t.Fatalf("test case [ %s ], wantArgFlagNotErr %v, but p.Exec() not error", tc.name, tc.wantArgFlagNotErr)
				}
				t.Logf("check args error: %v", errPluginRun)
			}
		})
	}
}

func TestPlugin(t *testing.T) {
	t.Log("do GiteaCCRelease")
	if envCheck(t) {
		return
	}
	if envMustArgsCheck(t) {
		return
	}
	t.Log("mock GiteaCCRelease")

	t.Log("mock gitea_cc_plugin config")

	testDataPathRoot, errTestDataPathRoot := testGoldenKit.GetOrCreateTestDataFullPath("plugin_test")
	if errTestDataPathRoot != nil {
		t.Fatal(errTestDataPathRoot)
	}

	// tagPipeline
	tagPipelineWoodpeckerInfo := *wd_mock.NewWoodpeckerInfo(
		wd_mock.FastWorkSpace(filepath.Join(testDataPathRoot, "tagPipeline")),
		wd_mock.FastTag("v1.0.0", "new tag"),
	)
	tagPipelineSettings := mockPluginSettings()

	tests := []struct {
		name           string
		woodpeckerInfo wd_info.WoodpeckerInfo
		settings       gitea_cc_plugin.Settings
		workRoot       string

		ossTransferKey  string
		ossTransferData interface{}

		isDryRun bool
		wantErr  bool
	}{
		{
			name:           "tagPipeline",
			woodpeckerInfo: tagPipelineWoodpeckerInfo,
			settings:       tagPipelineSettings,
			isDryRun:       true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := mockPluginWithSettings(t, tc.woodpeckerInfo, tc.settings)
			p.Settings.DryRun = tc.isDryRun
			if tc.ossTransferKey != "" {
				errGenTransferData := generateTransferStepsOut(
					p,
					tc.ossTransferKey,
					tc.ossTransferData,
				)
				if errGenTransferData != nil {
					t.Fatal(errGenTransferData)
				}
			}
			err := p.Exec()
			if (err != nil) != tc.wantErr {
				t.Errorf("FeishuPlugin.Exec() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
