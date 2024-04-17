# Changelog

All notable changes to this project will be documented in this file. See [convention-change-log](https://github.com/convention-change/convention-change-log) for commit guidelines.

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
