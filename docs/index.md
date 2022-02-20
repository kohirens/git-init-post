# Git Tool Belt

## Local Development

### Testing

This tool makes use of Git commands which requires a working on repository. For
testing we don't want to use a real repository. However, we do wish to test
and validate the commands are performing as expected. One
approach is to make Git bundles (archives) and un-bundle them during test runs.
It's one cool way to make test repos that work very well.

**How to make a fixture repository for test**

1. Make a new folder in `testdata`, you can follow the existing naming
   convention.
2. cd into that new directory and run `git init` to initialize it.
3. Now just add files and commit them in this directory. Be careful. Make Sure
   you do **NOT** commit any of these test repository files to the main project.
4. Once you get the test repo to a state that you want, it's time to bundle it
   up using the Git bundle command:
   ```
   git bundle create <bundle-filename> <branch>
   ```
5. So from inside the test repository directory run the command, for
   example
   ```
   git bundle create ../repo-04.bundle main
   ```
   In this example we save the bundle in the `testdata` directory. Please be
   sure to save yours there as well by using `../` before the name of the bundle.
6. Now go back to the root of the main project and be sure to commit the
   `*.bundle` file to the main project.

NOTE: There is a function that you can use in the test to un-bundle this file during
the test run, for example:
```go
tmpRepo := setupARepository("repo-04", "taggable-04")
```
tempRepo points to the path where the repo was extracted. Also know that this
function will append ".bundle" to the first parameter to find the actual bundle.