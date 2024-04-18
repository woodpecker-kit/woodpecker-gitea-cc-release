package gitea_cc_plugin

import (
	"code.gitea.io/sdk/gitea"
	"fmt"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"os"
	"path/filepath"
)

func (r *releaseClient) GetUploadDesc() string {
	return r.uploadDesc
}

func (r *releaseClient) Title() string {
	return r.title
}

func (r *releaseClient) SetTitle(title string) {
	r.title = title
}

func (r *releaseClient) Tag() string {
	return r.tag
}

func (r *releaseClient) SetNote(noteContent string) {
	r.note = noteContent
}

// BuildRelease retrieves or creates a release
func (r *releaseClient) BuildRelease() (*gitea.Release, error) {
	release, err := r.getRelease()
	if err != nil && release == nil {
		wd_log.Debugf("not getRelease release but can try new release, err: %v", err)
	} else if release != nil {
		wd_log.Infof("gitea release found Release ID:%d Draft:%v Prerelease:%v url: %s", release.ID, release.IsDraft, release.IsPrerelease, release.HTMLURL)
		return release, nil
	}

	if r.dryRun {
		wd_log.Infof("~> dry run mode open, not creating release\n")
		wd_log.Infof("-> try to create release %s/%s/%s\n", r.GetBaseUrl(), r.owner, r.repo)
		return &gitea.Release{
			ID: -1,
		}, nil
	}

	// if no release was found by that tag, create a new one

	release, err = r.newRelease()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve or create a release: %s", err)
	}

	wd_log.Infof("new Release ID:%d Draft:%v Prerelease:%v url: %s", release.ID, release.IsDraft, release.IsPrerelease, release.HTMLURL)
	return release, nil
}

// UploadFiles uploads files to a release
func (r *releaseClient) UploadFiles(releaseID int64) error {
	if len(r.uploadFilePaths) == 0 {
		if r.uploadDesc != "" {
			wd_log.Infof("gitea release client upload files desc:\n%s\n", r.uploadDesc)
			return nil
		}
		wd_log.Infof("check settings no upload files found!\n")
		return nil
	}

	if r.dryRun {
		wd_log.Infof("try to upload files to release %s/%s\n", r.owner, r.repo)
		for _, filePath := range r.uploadFilePaths {
			wd_log.Infof("-> try upload file: %s\n", filePath)
		}
		wd_log.Infof("~> dry run, not uploading files\n")
		return nil
	}

	attachments, _, err := r.GiteaClient().ListReleaseAttachments(r.owner, r.repo, releaseID, gitea.ListReleaseAttachmentsOptions{})
	if err != nil {
		return fmt.Errorf("failed to fetch existing assets: %s", err)
	}

	var uploadFiles []string

files:

	for _, filePath := range r.uploadFilePaths {
		for _, attachment := range attachments {
			if attachment.Name == filepath.Base(filePath) {
				switch r.fileExistsDo {
				case FileExistsDoOverwrite:
					// do nothing now we will delete the old file and upload the new one
				case FileExistsDoFail:
					return fmt.Errorf("asset file %s already exists", filepath.Base(filePath))
				case FileExistsDoSkip:
					wd_log.Infof("skipping pre-existing %s artifact\n", attachment.Name)
					continue files
				default:
					return fmt.Errorf("internal error, unkown file_exist value %s", r.fileExistsDo)
				}
			}
		}

		uploadFiles = append(uploadFiles, filePath)
	}

	for _, file := range uploadFiles {
		handle, errOpen := os.Open(file)

		if errOpen != nil {
			return fmt.Errorf("failed to read %s artifact: %s", file, errOpen)
		}

		fileBaseName := filepath.Base(file)

		for _, attachment := range attachments {
			if attachment.Name == fileBaseName {
				if _, err := r.GiteaClient().DeleteReleaseAttachment(r.owner, r.repo, releaseID, attachment.ID); err != nil {
					return fmt.Errorf("failed to delete file base name: %s artifact: %s", fileBaseName, err)
				}

				wd_log.Infof("successfully deleted old attachment.ID[ %v ] artifact %s\n", attachment.ID, attachment.Name)
			}
		}

		if _, _, err = r.GiteaClient().CreateReleaseAttachment(r.owner, r.repo, releaseID, handle, fileBaseName); err != nil {
			return fmt.Errorf("failed to upload file base name: %s artifact: %s", fileBaseName, err)
		}

		wd_log.Infof("successfully uploaded artifact file name [ %s ] path: %s \n", fileBaseName, file)
	}

	return nil
}

func (r *releaseClient) getRelease() (*gitea.Release, error) {
	releases, _, err := r.GiteaClient().ListReleases(r.owner, r.repo, gitea.ListReleasesOptions{})
	if err != nil {
		return nil, err
	}
	for _, release := range releases {
		if release.TagName == r.tag {
			wd_log.Debugf("Successfully retrieved %s release\n", r.tag)
			return release, nil
		}
	}
	return nil, fmt.Errorf("release %s not found", r.tag)
}

func (r *releaseClient) newRelease() (*gitea.Release, error) {
	c := gitea.CreateReleaseOption{
		TagName:      r.tag,
		Target:       r.tagTarget,
		IsDraft:      r.draft,
		IsPrerelease: r.prerelease,
		Title:        r.title,
		Note:         r.note,
	}

	release, _, err := r.GiteaClient().CreateRelease(r.owner, r.repo, c)
	if err != nil {
		return nil, fmt.Errorf("failed to create release: %s", err)
	}

	return release, nil
}
