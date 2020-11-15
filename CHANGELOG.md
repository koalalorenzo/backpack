# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added 
### Removed
### Changed
### Fixed

## [0.2.0] - 2020-11-15
### Added 
- It is possible to plan changes and check the differences before applying them via `backpack plan` ([!11](https://gitlab.com/Qm64/backpack/-/merge_requests/11), [#12](https://gitlab.com/Qm64/backpack/-/issues/12))
- It is possible to stop jobs of a pack using `backpack stop` ([!13](https://gitlab.com/Qm64/backpack/-/merge_requests/13), [#15](https://gitlab.com/Qm64/backpack/-/issues/15))
- CLI Aliases for Helm users (ex: `backpack uninstall` is `backpack stop`)([!13](https://gitlab.com/Qm64/backpack/-/merge_requests/13))
- It is possible to check the jobs' allocations of a pack using `backpack status` ([!14](https://gitlab.com/Qm64/backpack/-/merge_requests/14))

### Changed
- Optimize the Connection struct and avoid repeating tasks ([!11](https://gitlab.com/Qm64/backpack/-/merge_requests/11))
- Refactoring of `backpack run` to share part of the code with `backpack plan` ([!11](https://gitlab.com/Qm64/backpack/-/merge_requests/11))
- Renamed type Backpack to Pack to differ the software from the packages ([!12](https://gitlab.com/Qm64/backpack/-/merge_requests/12)))
- Upgraded to latest version of the dependencies ([!12](https://gitlab.com/Qm64/backpack/-/merge_requests/12)))

### Fixed
- Fixes network selection on `redis` and `fabiolb` examples when running docker ([!12](https://gitlab.com/Qm64/backpack/-/merge_requests/12)))

## [0.1.0] - 2020-10-12
### Added
This is the first release!
- Backpack now can to pack and unpack packages/bundles ([#2](https://gitlab.com/Qm64/backpack/-/issues/2))
- Backpack can help creating new packages from example ([#9](https://gitlab.com/Qm64/backpack/-/issues/9))
- Backpack can create new Hashicorp Nomad Jobs from a package ([#3](https://gitlab.com/Qm64/backpack/-/issues/3))
- The CLI accepts both URLs and local paths ([#6](https://gitlab.com/Qm64/backpack/-/issues/6))
- There are few example backpack files ([#5](https://gitlab.com/Qm64/backpack/-/issues/5))
- A backpack includes markdown documentation ([#8](https://gitlab.com/Qm64/backpack/-/issues/8))
- Ability to know what version of backpack I am running ([!9](https://gitlab.com/Qm64/backpack/-/merge_requests/9))
