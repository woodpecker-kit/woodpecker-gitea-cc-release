# Changelog

All notable changes to this project will be documented in this file. See [convention-change-log](https://github.com/convention-change/convention-change-log) for commit guidelines.

## [1.2.3](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/compare/1.2.2...v1.2.3) (2024-09-29)

### 👷‍ Build System

* bump github.com/convention-change/convention-change-log ([569eef8e](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/569eef8efd745f17305d7800ef587a806e1cd1e6))

* bump github.com/sinlov-go/unittest-kit from 1.1.1 to 1.2.1 ([b4232ce3](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/b4232ce34c03378d02fbe8ce6b30e2e28fd32ad0))

* bump golang.org/x/crypto from 0.26.0 to 0.27.0 ([0061d310](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/0061d31097d83732fd168457d091cad7756b80b4))

* bump github.com/sinlov-go/go-common-lib from 1.7.0 to 1.7.1 ([15780a67](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/15780a679b50d74cd8462a71ceb69cb56f6e0719))

* bump github.com/convention-change/convention-change-log ([693c6c9f](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/693c6c9f197e74c382370b0e857e856ec08e3195))

* bump github.com/Masterminds/semver/v3 from 3.2.1 to 3.3.0 ([00293a08](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/00293a0892ad0cb9fe545fa717277b05208894d4))

* bump github.com/sinlov-go/unittest-kit from 1.1.0 to 1.1.1 ([c4f6b10f](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/c4f6b10f8c93d83c04bb0e4116e274e890e47b30))

* bump github.com/urfave/cli/v2 from 2.27.3 to 2.27.4 ([a1a9ea39](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/a1a9ea3999c71064c8c3f8ecbebff8dc3c11ff17))

* bump golang.org/x/crypto from 0.25.0 to 0.26.0 ([3a0fe562](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/3a0fe5622fab488f32c48d6ac2dc176eb93faf27))

* change go version to 1.21.11 to check build and run ([a4895d44](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/a4895d44d6f46786f5159b53b13583d91f924733)), fix [#29](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/issues/29)

* add auto-merge-dependabot.yml and let docker-buildx bake plugin fast ([dc4e18f3](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/dc4e18f35d128b61d0e972cd16444290248f4b24))

## [1.2.2](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/compare/1.2.1...v1.2.2) (2024-04-23)

### 🐛 Bug Fixes

* flag `settngs.gitea-timeout-second` not effective ([358856bc](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/358856bc1db6cac330dc356b84f6d9946c582796)), fix [#15](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/issues/15)

## [1.2.1](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/compare/1.2.0...v1.2.1) (2024-04-21)

## [1.2.0](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/compare/1.1.2...v1.2.0) (2024-04-19)

### ✨ Features

* docker platform `linux/amd64 linux/386 linux/arm64/v8 linux/arm/v7 linux/ppc64le linux/s390x` ([701460ef](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/701460ef896a5ff2725ca212428ffad0583fcfbf)), feat [#11](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/issues/11)

## [1.1.2](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/compare/1.1.1...v1.1.2) (2024-04-18)

### 🐛 Bug Fixes

* fix file check sum error ([4a1e9ebf](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/4a1e9ebfbed49afd36aa6bac484d9901a3664bdb)), fix [#9](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/issues/9)

## [1.1.1](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/compare/1.1.0...v1.1.1) (2024-04-18)

## [1.1.0](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/compare/1.0.0...v1.1.0) (2024-04-18)

### ✨ Features

* `github.com/sinlov-go/gitea-client-wrapper` pkg for use GiteaTokenClient to management ([bb72c317](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/bb72c31771c11850c5f1527e2cf06c4fc176d616))

### 📝 Documentation

* update doc of usage ([7541fd3b](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/7541fd3b31a89e5f487aadbd55bc7dea8aea20a2))

## 1.0.0 (2024-04-16)

### ✨ Features

* when `CI_FORGE_TYPE` is [ gitea ], and `gitea-base-url` is empty, will try get `CI_FORGE_URL` ([21e33e5a](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/21e33e5a7582f8c05d33a71be393a40f93438ef4))

* let gitea client timeout be set ([1f06d1b1](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/1f06d1b1ee914328376ed0802897dbb750496b63))

* add flag `gitea-release-note-by-convention-change` and more config to make release ([da5193c1](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/da5193c1d9e28cb48b2bfddf3241f9cf81e661aa))

* add flag for relase ready and flag args check test case ([2f03147d](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/2f03147d1bf15fa8fd867a3efd7c4c6fb67910c9))

### 📝 Documentation

* add useage doc/docs.md and remove check at gitea-dry-run mode event ([07d1727a](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/07d1727adc1c1e689112cdd232f3afac61bf7db1))

### ♻ Refactor

* move pkg to `gitea_cc_plugin` ([26a94e2b](https://github.com/woodpecker-kit/woodpecker-gitea-cc-release/commit/26a94e2b60beb09daeb4d105bc77115adefbada6))
