package gitea_cc_plugin

import (
	"github.com/convention-change/convention-change-log/changelog"
	"github.com/convention-change/convention-change-log/convention"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"path/filepath"
)

func (p *GiteaCCRelease) releaseByClient() error {
	rc, errNewReleaseClient := NewReleaseClientByWoodpeckerShort(p.ShortInfo(), p.Settings)
	if errNewReleaseClient != nil {
		return errNewReleaseClient
	}

	if p.Settings.GiteaReleaseNoteByConventionChange {

		specFilePath := filepath.Join(p.Settings.RootPath, versionRcFileName)
		changeLogSpec, errChangeLogSpecByPath := convention.LoadConventionalChangeLogSpecByPath(specFilePath)
		if errChangeLogSpecByPath != nil {
			wd_log.Error(errChangeLogSpecByPath)
			return errChangeLogSpecByPath
		}
		reader, errCC := changelog.NewReader(p.Settings.GiteaReleaseConventionReadPath, *changeLogSpec)
		if errCC == nil {
			rc.SetNote(reader.HistoryFirstContent())
			rc.SetTitle(reader.HistoryFirstTagShort())
		} else {
			wd_log.Warnf("not found change log or other error: %v\n", errCC)
		}
	}

	release, errBuildRelease := rc.BuildRelease()
	if errBuildRelease != nil {
		wd_log.Error(errBuildRelease)
		return errBuildRelease
	}

	if errUpload := rc.UploadFiles(release.ID); errUpload != nil {
		wd_log.Error(errUpload)
		return errUpload
	}

	return nil
}
