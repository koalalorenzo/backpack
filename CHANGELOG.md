# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added 
- It is possible to plan changes and check the differences before applying them via `backpack plan` ([mr11](https://gitlab.com/Qm64/backpack/-/merge_requests/11), [#12](https://gitlab.com/Qm64/backpack/-/issues/12))
### Removed
### Changed
- Optimize the Connection struct and avoid repeating tasks ([mr11](https://gitlab.com/Qm64/backpack/-/merge_requests/11))
- Refactoring of `backpack run` to share part of the code with `backpack plan` ([mr11](https://gitlab.com/Qm64/backpack/-/merge_requests/11))
- Renamed type Backpack to Pack to differ the software from the packages ([mr12](https://gitlab.com/Qm64/backpack/-/merge_requests/12)))
- Upgraded to latest version of the dependencies ([mr12](https://gitlab.com/Qm64/backpack/-/merge_requests/12)))
### Fixed
- Fixes network selection on `redis` and `fabiolb` examples when running docker ([mr12](https://gitlab.com/Qm64/backpack/-/merge_requests/12)))

## [1.0.0] - 2020-10-12
### Added
This is the first release!
- Backpack now can to pack and unpack packages/bundles ([#2](https://gitlab.com/Qm64/backpack/-/issues/2))
- Backpack can help creating new packages from example ([#9](https://gitlab.com/Qm64/backpack/-/issues/9))
- Backpack can create new Hashicorp Nomad Jobs from a package ([#3](https://gitlab.com/Qm64/backpack/-/issues/3))
- The CLI accepts both URLs and local paths ([#6](https://gitlab.com/Qm64/backpack/-/issues/6))
- There are few example backpack files ([#5](https://gitlab.com/Qm64/backpack/-/issues/5))
- A backpack includes markdown documentation ([#8](https://gitlab.com/Qm64/backpack/-/issues/8))
- Ability to know what version of backpack I am running ([mr9](https://gitlab.com/Qm64/backpack/-/merge_requests/9))
