# TODO

## Common task

* Add pre commit git hook to this project to perform the following:
    * `go fmt` - to format the project
    * `git-chglog --output CHANGELOG.md` - to update the changelog
    * then append the changes then commit.
* Output build info to terminal
* Rename sub-command `version` to something else.
* Output help for all global and sub-commands.
* Build and push docker image in CI
* Build Linux executable and upload to Github.

## Feature Level

* Figure out how clients can add a way to include the command to generate
  build version info to their `go builds`.

* Generate a pre commit git hook to perform the following:
  * `git-chglog --output CHANGELOG.md` - should it produce output, then fail to
    commit and display what should be added to the CHANGELOG.
