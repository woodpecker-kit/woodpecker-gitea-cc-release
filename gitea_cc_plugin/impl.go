package gitea_cc_plugin

import (
	"fmt"
	"github.com/sinlov-go/go-common-lib/pkg/string_tools"
	"github.com/sinlov-go/go-common-lib/pkg/struct_kit"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
)

func (p *GiteaCCRelease) ShortInfo() wd_short_info.WoodpeckerInfoShort {
	if p.wdShortInfo == nil {
		info2Short := wd_short_info.ParseWoodpeckerInfo2Short(*p.woodpeckerInfo)
		p.wdShortInfo = &info2Short
	}
	return *p.wdShortInfo
}

// SetWoodpeckerInfo
// also change ShortInfo() return
func (p *GiteaCCRelease) SetWoodpeckerInfo(info wd_info.WoodpeckerInfo) {
	var newInfo wd_info.WoodpeckerInfo
	_ = struct_kit.DeepCopyByGob(&info, &newInfo)
	p.woodpeckerInfo = &newInfo
	info2Short := wd_short_info.ParseWoodpeckerInfo2Short(newInfo)
	p.wdShortInfo = &info2Short
}

func (p *GiteaCCRelease) GetWoodPeckerInfo() wd_info.WoodpeckerInfo {
	return *p.woodpeckerInfo
}

func (p *GiteaCCRelease) OnlyArgsCheck() {
	p.onlyArgsCheck = true
}

func (p *GiteaCCRelease) Exec() error {
	errLoadStepsTransfer := p.loadStepsTransfer()
	if errLoadStepsTransfer != nil {
		return errLoadStepsTransfer
	}

	errCheckArgs := p.checkArgs()
	if errCheckArgs != nil {
		return fmt.Errorf("check args err: %v", errCheckArgs)
	}

	if p.onlyArgsCheck {
		wd_log.Info("only check args, skip do doBiz")
		return nil
	}

	err := p.doBiz()
	if err != nil {
		return err
	}
	errSaveStepsTransfer := p.saveStepsTransfer()
	if errSaveStepsTransfer != nil {
		return errSaveStepsTransfer
	}

	return nil
}

func (p *GiteaCCRelease) loadStepsTransfer() error {
	return nil
}

func (p *GiteaCCRelease) checkArgs() error {

	if p.Settings.GiteaBaseUrl == "" {
		return fmt.Errorf("check args [ %s ] must set, now is empty", CliNameGiteaBaseUrl)
	}
	if p.Settings.GiteaApiKey == "" {
		return fmt.Errorf("check args [ %s ] must set, now is empty", CliNameGiteaApiKey)
	}

	if len(p.Settings.GiteaReleaseFilesGlobs) > 0 {
		errFileExistsDoList := argCheckInArr(CliNameGiteaReleaseFileExistsDo, p.Settings.GiteaReleaseFileExistsDo, supportFileExistsDoList)
		if errFileExistsDoList != nil {
			return errFileExistsDoList
		}
		if len(p.Settings.GiteaFilesChecksum) > 0 {
			for _, checkCfg := range p.Settings.GiteaFilesChecksum {
				errCheckSumSupport := argCheckInArr(CliNameGiteaFilesChecksum, checkCfg, CheckSumSupport)
				if errCheckSumSupport != nil {
					return errCheckSumSupport
				}
			}
		}
	}

	return nil
}

func argCheckInArr(mark string, target string, checkArr []string) error {
	if !(string_tools.StringInArr(target, checkArr)) {
		return fmt.Errorf("not support %s now [ %s ], must in %v", mark, target, checkArr)
	}
	return nil
}

// doBiz
//
//	replace this code with your gitea_cc_plugin implementation
func (p *GiteaCCRelease) doBiz() error {

	err := p.releaseByClient()
	if err != nil {
		return err
	}

	return nil
}

func (p *GiteaCCRelease) saveStepsTransfer() error {
	// remove or change this code

	if p.Settings.StepsOutDisable {
		wd_log.Debugf("steps out disable by flag [ %v ], skip save steps transfer", p.Settings.StepsOutDisable)
		return nil
	}

	return nil
}
