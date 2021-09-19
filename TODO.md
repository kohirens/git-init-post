Common task:
Setup a changelog
add a way for the version to be included when go build is run. default to use the git hash, and use a flag to set from the git tag.
after git commit run hooks like:
go fmt - to format the project
git-chglog - to update the changelog 