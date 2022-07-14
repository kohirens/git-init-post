<a name="unreleased"></a>
## [Unreleased]


<a name="2.0.1"></a>
## [2.0.1] - 2022-07-14
### Fixed
- checking nil value.


<a name="2.0.0"></a>
## [2.0.0] - 2022-06-30
### Added
- rel tag can now have context after semantic version.

### Changed
- Renamed sub-command checkConf to checkconf.
- Renamed method from check to validate for flags.


<a name="1.2.9"></a>
## [1.2.9] - 2022-05-16
### Fixed
- Building Executables.


<a name="1.2.8"></a>
## [1.2.8] - 2022-05-16
### Removed
- Incorrect flag on code generator.


<a name="1.2.7"></a>
## [1.2.7] - 2022-05-16
### Fixed
- Type in template.


<a name="1.2.6"></a>
## [1.2.6] - 2022-05-16
### Fixed
- Logic in CI workflows


<a name="1.2.5"></a>
## [1.2.5] - 2022-05-16
### Fixed
- Publishing Artifacts


<a name="1.2.4"></a>
## [1.2.4] - 2022-05-16
### Changed
- Use Kohirens circleci-go image to publish executables.

### Fixed
- Auto-publish, upgraded to unreleased Version Release Orb.
- varName flag not passed in for semver subcommand.


<a name="1.2.3"></a>
## [1.2.3] - 2022-04-22

<a name="1.2.2"></a>
## [1.2.2] - 2022-04-21

<a name="1.2.1"></a>
## [1.2.1] - 2022-04-21

<a name="1.2.0"></a>
## [1.2.0] - 2022-04-21

<a name="1.1.3"></a>
## [1.1.3] - 2022-03-08
### Fixed
- Version info.


<a name="1.1.2"></a>
## [1.1.2] - 2022-03-07

<a name="1.1.1"></a>
## [1.1.1] - 2022-03-06
### Changed
- Removed hard-code version from publising exe's in CI.


<a name="1.1.0"></a>
## [1.1.0] - 2022-03-06
### Changed
- Publish binaries.


<a name="1.0.7"></a>
## [1.0.7] - 2022-02-22
### Removed
- VS code server download.


<a name="1.0.6"></a>
## [1.0.6] - 2022-02-22
### Fixed
- Docker build in CI.


<a name="1.0.5"></a>
## [1.0.5] - 2022-02-22
### Fixed
- Docker build permission error.


<a name="1.0.4"></a>
## [1.0.4] - 2022-02-22
### Fixed
- Docker build permission error.


<a name="1.0.3"></a>
## [1.0.3] - 2022-02-22
### Fixed
- Docker build permission error.


<a name="1.0.2"></a>
## [1.0.2] - 2022-02-22
### Fixed
- Docker build permission error.


<a name="1.0.1"></a>
## [1.0.1] - 2022-02-22
### Fixed
- Missing checkout step when publishing image.


<a name="1.0.0"></a>
## [1.0.0] - 2022-02-22
### Changed
- Build and publish docker image in CI.


<a name="0.9.0"></a>
## [0.9.0] - 2022-02-20

<a name="0.2.1"></a>
## [0.2.1] - 2022-02-20
### Changed
- Allow other revision ranges for taggable.
- Exit gracefully when no subcommand is passed in.


<a name="0.2.0"></a>
## [0.2.0] - 2022-01-31
### Added
- Verbose flag for taggable subcommand.
- Version Sub Comand repo Flag
- Taggable Sub Command
- Check if commit is taggable.
- jq tool.
- Get the next release version from the git logs.
- Get next version (basic).
- Method to add commit hash to build version info.
- CHANGELOG.md.

### Changed
- Zero out minor and patch verision with major and minor change.
- Every user in the container can access the tool.
- Refactored to introduce a programming API.

### Removed
- Unused import from helpers.


<a name="0.1.0"></a>
## 0.1.0 - 2021-09-19
### Added
- Version subcommand.


[Unreleased]: https://github.com/kohirens/git-tool-belt/compare/2.0.1...HEAD
[2.0.1]: https://github.com/kohirens/git-tool-belt/compare/2.0.0...2.0.1
[2.0.0]: https://github.com/kohirens/git-tool-belt/compare/1.2.9...2.0.0
[1.2.9]: https://github.com/kohirens/git-tool-belt/compare/1.2.8...1.2.9
[1.2.8]: https://github.com/kohirens/git-tool-belt/compare/1.2.7...1.2.8
[1.2.7]: https://github.com/kohirens/git-tool-belt/compare/1.2.6...1.2.7
[1.2.6]: https://github.com/kohirens/git-tool-belt/compare/1.2.5...1.2.6
[1.2.5]: https://github.com/kohirens/git-tool-belt/compare/1.2.4...1.2.5
[1.2.4]: https://github.com/kohirens/git-tool-belt/compare/1.2.3...1.2.4
[1.2.3]: https://github.com/kohirens/git-tool-belt/compare/1.2.2...1.2.3
[1.2.2]: https://github.com/kohirens/git-tool-belt/compare/1.2.1...1.2.2
[1.2.1]: https://github.com/kohirens/git-tool-belt/compare/1.2.0...1.2.1
[1.2.0]: https://github.com/kohirens/git-tool-belt/compare/1.1.3...1.2.0
[1.1.3]: https://github.com/kohirens/git-tool-belt/compare/1.1.2...1.1.3
[1.1.2]: https://github.com/kohirens/git-tool-belt/compare/1.1.1...1.1.2
[1.1.1]: https://github.com/kohirens/git-tool-belt/compare/1.1.0...1.1.1
[1.1.0]: https://github.com/kohirens/git-tool-belt/compare/1.0.7...1.1.0
[1.0.7]: https://github.com/kohirens/git-tool-belt/compare/1.0.6...1.0.7
[1.0.6]: https://github.com/kohirens/git-tool-belt/compare/1.0.5...1.0.6
[1.0.5]: https://github.com/kohirens/git-tool-belt/compare/1.0.4...1.0.5
[1.0.4]: https://github.com/kohirens/git-tool-belt/compare/1.0.3...1.0.4
[1.0.3]: https://github.com/kohirens/git-tool-belt/compare/1.0.2...1.0.3
[1.0.2]: https://github.com/kohirens/git-tool-belt/compare/1.0.1...1.0.2
[1.0.1]: https://github.com/kohirens/git-tool-belt/compare/1.0.0...1.0.1
[1.0.0]: https://github.com/kohirens/git-tool-belt/compare/0.9.0...1.0.0
[0.9.0]: https://github.com/kohirens/git-tool-belt/compare/0.2.1...0.9.0
[0.2.1]: https://github.com/kohirens/git-tool-belt/compare/0.2.0...0.2.1
[0.2.0]: https://github.com/kohirens/git-tool-belt/compare/0.1.0...0.2.0
