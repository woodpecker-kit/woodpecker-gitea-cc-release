---
name: woodpecker-gitea-cc-release
description: woodpecker gitea_cc_plugin template
author: woodpecker-kit
tags: [ environment, woodpecker-gitea-cc-release ]
containerImage: woodpecker-kit/woodpecker-gitea-cc-release
containerImageUrl: https://hub.docker.com/r/woodpecker-kit/woodpecker-gitea-cc-release
url: https://github.com/woodpecker-kit/woodpecker-gitea-cc-release
icon: https://raw.githubusercontent.com/woodpecker-kit/woodpecker-gitea-cc-release/main/doc/logo.svg
---

woodpecker plugin template

## Settings

| Name    | Required | Default value | Description                                  |
|---------|----------|---------------|----------------------------------------------|
| `debug` | **no**   | *false*       | open debug log or open by env `PLUGIN_DEBUG` |

**Hide Settings:**

| Name                                        | Required | Default value                    | Description                                                                      |
|---------------------------------------------|----------|----------------------------------|----------------------------------------------------------------------------------|
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
    image: sinlov/woodpecker-gitea-cc-release:latest
    pull: false
    settings:
    # debug: true
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
```

## Notes

- Please add notes

## Known limitations

- Please add a known issue
