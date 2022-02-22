<a name="unreleased"></a>
## [Unreleased]


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


[Unreleased]: https://github.com/kohirens/git-tool-belt/compare/1.0.0...HEAD
[1.0.0]: https://github.com/kohirens/git-tool-belt/compare/0.9.0...1.0.0
[0.9.0]: https://github.com/kohirens/git-tool-belt/compare/0.2.1...0.9.0
[0.2.1]: https://github.com/kohirens/git-tool-belt/compare/0.2.0...0.2.1
[0.2.0]: https://github.com/kohirens/git-tool-belt/compare/0.1.0...0.2.0
