package gitea_cc_plugin

import (
	"code.gitea.io/sdk/gitea"
	"fmt"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"os"
	"path"
	"path/filepath"
)

func (r *releaseClient) GetUploadDesc() string {
	return r.uploadDesc
}

// SetOTP sets OTP for 2FA
func (r *releaseClient) SetOTP(otp string) {
	r.mutex.Lock()
	r.otp = otp
	r.client.SetOTP(otp)
	r.mutex.Unlock()
}

// SetSudo sets username to impersonate.
func (r *releaseClient) SetSudo(sudo string) {
	r.mutex.Lock()
	r.sudo = sudo
	r.client.SetSudo(sudo)
	r.mutex.Unlock()
}

// SetBasicAuth sets username and password
func (r *releaseClient) SetBasicAuth(username, password string) {
	r.mutex.Lock()
	r.username, r.password = username, password
	r.client.SetBasicAuth(username, password)
	r.mutex.Unlock()
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
		wd_log.Infof("try to create release %s/%s\n", r.owner, r.repo)
		wd_log.Infof("dry run, not creating release\n")
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
		wd_log.Infof("dry run, not uploading files\n")
		return nil
	}

	attachments, _, err := r.client.ListReleaseAttachments(r.owner, r.repo, releaseID, gitea.ListReleaseAttachmentsOptions{})
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
					// do nothing
				case FileExistsDoFail:
					return fmt.Errorf("asset file %s already exists", path.Base(filePath))
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

		for _, attachment := range attachments {
			if attachment.Name == path.Base(file) {
				if _, err := r.client.DeleteReleaseAttachment(r.owner, r.repo, releaseID, attachment.ID); err != nil {
					return fmt.Errorf("failed to delete %s artifact: %s", file, err)
				}

				wd_log.Infof("successfully deleted old attachment.ID[ %v ] artifact %s\n", attachment.ID, attachment.Name)
			}
		}

		if _, _, err = r.client.CreateReleaseAttachment(r.owner, r.repo, releaseID, handle, path.Base(file)); err != nil {
			return fmt.Errorf("failed to upload %s artifact: %s", file, err)
		}

		wd_log.Infof("successfully uploaded artifact: %s \n", file)
	}

	return nil
}

func (r *releaseClient) getRelease() (*gitea.Release, error) {
	releases, _, err := r.client.ListReleases(r.owner, r.repo, gitea.ListReleasesOptions{})
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

	release, _, err := r.client.CreateRelease(r.owner, r.repo, c)
	if err != nil {
		return nil, fmt.Errorf("failed to create release: %s", err)
	}

	return release, nil
}
