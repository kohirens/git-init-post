# Kohirens: Release  Tools

A collection of tools to help you release software.

[![CircleCI](https://dl.circleci.com/status-badge/img/gh/kohirens/git-tool-belt/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/kohirens/git-tool-belt/tree/main)

## Features

## Sub Command "semver"

Generate a current and next tag info based on special tags in the commit messages.
Currently, special tags refer to commit message which begin with a tag of:

| Tag              | Description                                                                                                                                                                |
|------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `add: `          | Marks that a new feature was added and increments the minor version.                                                                                                       |
| `chg: `          | A standard change and increments the patch version.                                                                                                                        |
| `dep: `          | Indicates a feature is deprecated and will increment the patch number.                                                                                                     |
| `rmv: `          | Marks a feature removed but not a breaking change, increments the patch number.                                                                                            |
| `rel: x.x.x-rc1` | Will use the version number specified after the tag, anything after the `-` is optional. The dash is important otherwise anything after the third digit will be truncated. |
| BREAKING CHANGE  | Will cause the major number to increment by 1.                                                                                                                             |

### Examples

`git-tool-belt semver`

or

`git-tool-belt semver -repo my/git/repo`

**Use Cases**

1. generate the file before a build to incorporate build info into your build artifacts.
2. Generate a tag for automating a release in your CI (Continuous Integration) pipeline.

## Sub Command "taggable"

Return "true" or "false" if commits in the given commit range contain any of the special tags. 

`git-tool-belt taggable`

or 

`git-tool-belt taggable -repo my/git/repo`

or

`git-tool-belt taggable -repo my/git/repo -commitRange 0.1.0..HEAD`

**Use Cases**

Check if there are changes in your commit worth incrementing the version number.
before trying to publish a release in a CI pipeline.
