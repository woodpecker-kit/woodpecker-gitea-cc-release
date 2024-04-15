---
name: woodpecker-gitea-cc-release
description: plugin for https://about.gitea.com/ and conventional commits logs, and files release
author: woodpecker-kit
tags: [ Gitea, publish ]
containerImage: sinlov/woodpecker-gitea-cc-release
containerImageUrl: https://hub.docker.com/r/sinlov/woodpecker-gitea-cc-release
url: https://github.com/woodpecker-kit/woodpecker-gitea-cc-release
icon: https://raw.githubusercontent.com/woodpecker-kit/woodpecker-gitea-cc-release/main/doc/logo.svg
---

woodpecker plugin for https://about.gitea.com/ and conventional commits logs to make release

## Settings

| Name                                      | Required | Default value  | Description                                                                                                      |
|-------------------------------------------|----------|----------------|------------------------------------------------------------------------------------------------------------------|
| `debug`                                   | **no**   | *false*        | open debug log or open by env `PLUGIN_DEBUG`                                                                     |
| `gitea-api-key`                           | **yes**  |                | gitea api key                                                                                                    |
| `gitea-base-url`                          | **no**   | ""             | when `CI_FORGE_TYPE` is `gitea`, and this flag is empty, will try get from env `CI_FORGE_URL`                    |
| `gitea-insecure`                          | **no**   | *false*        | gitea insecure enable                                                                                            |
| `gitea-dry-run`                           | **no**   | *false*        | dry run mode enable                                                                                              |
| `gitea-draft`                             | **no**   | *false*        | release draft enable                                                                                             |
| `gitea-prerelease`                        | **no**   | *true*         | prerelease enable                                                                                                |
| `gitea-release-file-root-path`            | **no**   | ""             | release file root path, if empty will use workspace, most of use project root path                               |
| `gitea-release-files-globs`               | **no**   |                | release as files by glob pattern, if empty will skip release files                                               |
| `gitea-release-file-exists-do`            | **no**   | *fail*         | do if release update file already exist, support: `[fail overwrite skip]` (default: "fail")                      |
| `gitea-files-checksum`                    | **no**   |                | generate specific checksums, empty will skip, support: `[md5 sha1 sha256 sha512 adler32 crc32 blake2b blake2s]`  |
| `gitea-release-title`                     | **no**   | ""             | release title, if empty will use tag, can be cover by tag name of convention change log                          |
| `gitea-release-note-by-convention-change` | **no**   | *false*        | release note by convention change, if true will read change log file, and cover flag settings.gitea-release-note |
| `gitea-release-read-change-log-file`      | **no**   | "CHANGELOG.md" | release read change log file, if empty will use default CHANGELOG.md                                             |                                             

**Hide Settings:**

| Name                                        | Required | Default value                    | Description                                                                      |
|---------------------------------------------|----------|----------------------------------|----------------------------------------------------------------------------------|
| `gitea-timeout-second`                      | **no**   | *10*                             | gitea release api timeout second, default 60, less 30                            |
| `timeout_second`                            | **no**   | *10*                             | command timeout setting by second                                                |
| `woodpecker-kit-steps-transfer-file-path`   | **no**   | `.woodpecker_kit.steps.transfer` | Steps transfer file path, default by `wd_steps_transfer.DefaultKitStepsFileName` |
| `woodpecker-kit-steps-transfer-disable-out` | **no**   | *false*                          | Steps transfer write disable out                                                 |

## Example

- workflow with backend `docker`

[![docker hub version semver](https://img.shields.io/docker/v/sinlov/woodpecker-gitea-cc-release?sort=semver)](https://hub.docker.com/r/sinlov/woodpecker-gitea-cc-release/tags?page=1&ordering=last_updated)
[![docker hub image size](https://img.shields.io/docker/image-size/sinlov/woodpecker-gitea-cc-release)](https://hub.docker.com/r/sinlov/woodpecker-gitea-cc-release)
[![docker hub image pulls](https://img.shields.io/docker/pulls/sinlov/woodpecker-gitea-cc-release)](https://hub.docker.com/r/sinlov/woodpecker-gitea-cc-release/tags?page=1&ordering=last_updated)

```yml
labels:
  backend: docker
steps:
  woodpecker-gitea-cc-release:
    image: sinlov/woodpecker-gitea-cc-release:latest # https://hub.docker.com/r/sinlov/woodpecker-gitea-cc-release
    pull: false
    settings:
      # debug: true
      gitea-api-key:
        from_secret: gitea_api_key_release
      # gitea-base-url: "" # when `CI_FORGE_TYPE` is `gitea`, and this flag is empty, will try get from env `CI_FORGE_URL`
      gitea-dry-run: true # dry run mode enable
      gitea-draft: true # release draft enable
      gitea-prerelease: true # prerelease enable, default is true
      gitea-release-files-globs:
        - "doc/*.md"
      gitea-release-file-exists-do: "skip" # do if release update file already exist, support: [fail overwrite skip] (default: "fail")
      gitea-files-checksum: # generate specific checksums, empty will skip, support: [md5 sha1 sha256 sha512 adler32 crc32 blake2b blake2s]
        - "md5"
        - "sha256"
        - "sha512"
      # gitea-release-note-by-convention-change: true # 
      gitea-release-title: "dry run release" # release title, if empty will use tag, can be cover by tag name of convention change log
```

- workflow with backend `local`, must install at local and effective at evn `PATH`
    - can download by [github release](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/releases)
- install at ${GOPATH}/bin, latest

```bash
go install -a github.com/woodpecker-kit/woodpecker-gitea-cc-release/cmd/woodpecker-gitea-cc-release@latest
```

[![GitHub release)](https://img.shields.io/github/v/release/woodpecker-kit/woodpecker-gitea-cc-release)](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/releases)

- install at ${GOPATH}/bin, v1.0.0

```bash
go install -v github.com/woodpecker-kit/woodpecker-gitea-cc-release/cmd/woodpecker-gitea-cc-release@v1.0.0
```

```yml
labels:
  backend: local
steps:
  woodpecker-gitea-cc-release:
    image: woodpecker-gitea-cc-release
    settings:
      # debug: true
      gitea-api-key:
        from_secret: gitea_api_key_release
      # gitea-base-url: "" # when `CI_FORGE_TYPE` is `gitea`, and this flag is empty, will try get from env `CI_FORGE_URL`
      gitea-dry-run: true # dry run mode enable
      gitea-draft: true # release draft enable
      gitea-prerelease: true # prerelease enable, default is true
      gitea-release-files-globs:
        - "doc/*.md"
      gitea-release-file-exists-do: "skip" # do if release update file already exist, support: [fail overwrite skip] (default: "fail")
      gitea-files-checksum: # generate specific checksums, empty will skip, support: [md5 sha1 sha256 sha512 adler32 crc32 blake2b blake2s]
        - "md5"
        - "sha256"
        - "sha512"
      # gitea-release-note-by-convention-change: true # 
      gitea-release-title: "dry run release" # release title, if empty will use tag, can be cover by tag name of convention change log
```

- total example

```yml
labels:
  backend: docker
steps:
  woodpecker-gitea-cc-release:
    image: sinlov/woodpecker-gitea-cc-release:latest # https://hub.docker.com/r/sinlov/woodpecker-gitea-cc-release
    pull: false
    settings:
      # debug: true
      gitea-timeout-second: 120 # gitea release api timeout second, default 60, less 30
      gitea-api-key:
        from_secret: gitea_api_key_release
      gitea-base-url: "" # when `CI_FORGE_TYPE` is `gitea`, and this flag is empty, will try get from env `CI_FORGE_URL`
      # gitea-insecure: true # gitea insecure enable
      gitea-dry-run: true # dry run mode enable
      gitea-draft: true # release draft enable
      gitea-prerelease: true # prerelease enable, default is true
      # gitea-release-file-root-path: "" # release file root path, if empty will use workspace, most of use project root path
      gitea-release-files-globs:
        - "doc/*.md"
      gitea-release-file-exists-do: "skip" # do if release update file already exist, support: [fail overwrite skip] (default: "fail")
      gitea-files-checksum: # generate specific checksums, empty will skip, support: [md5 sha1 sha256 sha512 adler32 crc32 blake2b blake2s]
        - "md5"
        - "sha256"
        - "sha512"
      gitea-release-title: "dry run release" # release title, if empty will use tag, can be cover by tag name of convention change log
      # gitea-release-note-by-convention-change: true # release note by convention change, if true will read change log file, and cover flag settings.gitea-release-note 
      # gitea-release-read-change-log-file: CHANGELOG.md # release read change log file, if empty will use default CHANGELOG.md
```

## Notes

- gitea client token use [Access Token](https://docs.gitea.com/development/api-usage#authentication)

## Known limitations

- dry run mode will not create release, just print release info
- release available for event `tag`
