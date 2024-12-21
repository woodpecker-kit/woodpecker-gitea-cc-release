[![ci](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/workflows/ci/badge.svg)](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/woodpecker-kit/woodpecker-gitea-cc-release?label=go.mod)](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release)
[![GoDoc](https://godoc.org/github.com/woodpecker-kit/woodpecker-gitea-cc-release?status.png)](https://godoc.org/github.com/woodpecker-kit/woodpecker-gitea-cc-release)
[![goreportcard](https://goreportcard.com/badge/github.com/woodpecker-kit/woodpecker-gitea-cc-release)](https://goreportcard.com/report/github.com/woodpecker-kit/woodpecker-gitea-cc-release)

[![GitHub license](https://img.shields.io/github/license/woodpecker-kit/woodpecker-gitea-cc-release)](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release)
[![codecov](https://codecov.io/gh/woodpecker-kit/woodpecker-gitea-cc-release/branch/main/graph/badge.svg)](https://codecov.io/gh/woodpecker-kit/woodpecker-gitea-cc-release)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/woodpecker-kit/woodpecker-gitea-cc-release)](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/tags)
[![GitHub release)](https://img.shields.io/github/v/release/woodpecker-kit/woodpecker-gitea-cc-release)](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/releases)

## for what

- this project used to woodpecker plugin

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/woodpecker-kit/woodpecker-gitea-cc-release)](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## Features

- [x] support [gitea](https://gitea.io/) version `gitea >= 1.11`
- [x] dry-run mode
- [X] gitea client token use [Access Token](https://docs.gitea.com/development/api-usage#authentication)
- [X] upload release files by glob pattern
- [X] support generate check sum file by `md5 sha1 sha256 sha512 crc32 ...`
    - [x] check sum file support tools as `md5sum sha1sum sha256sum sha512sum`
    - [x] generate file base name dot not duplicates by base name
- [X] support [conventional-commits](https://www.conventionalcommits.org/) log from `CHANGELOG.md`
  - from lib [github.com/convention-change/convention-change-log](https://github.com/convention-change/convention-change-log)
  - the same support [www.contributor-covenant.org](https://www.contributor-covenant.org/) change logs
- [x] docker platform support (v1.2.+)
  -  `linux/amd64 linux/386 linux/arm64/v8 linux/arm/v7 linux/ppc64le linux/s390x` 
- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## usage

### workflow usage

- see [doc](doc/docs.md)

## Notice

- want dev this project, see [dev doc](doc/README.md)