# TODO

## Common task

* Add pre commit git hook to this project to perform the following:
    * `go fmt` - to format the project
    * `git-chglog --output CHANGELOG.md` - to update the changelog
    * then append the changes then commit.

## Feature Level

* Figure out how clients can add a way to include the command to generate
  build version info to their `go builds`.

* Generate a pre commit git hook to perform the following:
  * `git-chglog --output CHANGELOG.md` - should it produce output, then fail to
    commit and display what should be added to the CHANGELOG.