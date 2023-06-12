// Produce semantic version information about a Git repository in JSON format.
// The content of which can be useful in the build process for the purpose of
// using as the version details for the application.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kohirens/git-tool-belt/pkg/help"
	"html/template"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type buildVersion struct {
	CommitHash        string `json:"CommitHash"`
	CurrentVersion    string `json:"currentVersion"`
	NextVersion       string `json:"nextVersion"`
	NextVersionReason string `json:"nextVersionReason"`
}

func semverMain(af *applicationFlags) error {
	// Default to the current working directory, or set it from the flag.
	repoPath, _ := os.Getwd()
	if len(af.semver.repo) > 0 {
		repoPath = af.semver.repo
	}

	svInfo, err1 := GetSemverInfo(repoPath)
	if err1 != nil {
		return err1
	}

	var svBytes []byte
	if af.semver.format == "go" {
		fmt.Println("generating go code")
		var err2 error
		svBytes, err2 = formatForGo(svInfo, af.semver.packageName, af.semver.varName)
		if err2 != nil {
			return err2
		}
	} else {
		var err3 error
		svBytes, err3 = json.Marshal(svInfo)
		if err3 != nil {
			return fmt.Errorf("could not JSON encode build version info, reason: %v", err3.Error())
		}
	}

	if af.semver.save != "" {
		// Write the info to a JSON file.
		return saveSemverInfo(af.semver.save, svBytes)
	} else {
		fmt.Printf("%s", svBytes)
	}

	return nil
}

// GetSemverInfo build a JSON file with semver info.
func GetSemverInfo(repoPath string) (*buildVersion, error) {
	// Check the path exist.
	fileObj, err := os.Stat(repoPath)
	if err != nil || !fileObj.IsDir() {
		return nil, fmt.Errorf("repository path does not exists: %v", repoPath)
	}

	bvInfo := new(buildVersion)
	bvInfo.CurrentVersion = getCurrentVersion(repoPath)
	bvInfo.NextVersion, bvInfo.NextVersionReason = getNextVersion(repoPath, bvInfo.CurrentVersion)
	// Add commit hash.
	hash, err2 := getCommitHash(repoPath, bvInfo.CurrentVersion)
	if err2 != nil {
		return nil, err2
	}
	bvInfo.CommitHash = hash

	return bvInfo, nil
}

// FormatForGo Convert the JSON into Go code.
func formatForGo(svInfo *buildVersion, pn, vn string) ([]byte, error) {
	svData := &help.ByteBuf{}

	tmplStr := `package {{ .PackageName }}
// Code generated .* DO NOT EDIT.
func init() {
	// Predefine this struct in your code with the following members of type string.
	{{ .VarName }}.CommitHash = "{{ .CommitHash }}"
	{{ .VarName }}.CurrentVersion = "{{ .CurrentVersion }}"
}
`
	tmpl, err1 := template.New("sv").Parse(tmplStr)

	if err1 != nil {
		return nil, err1
	}

	fmt.Printf("generating code for package %v\n", pn)

	placeholders := struct {
		CommitHash, CurrentVersion, PackageName, VarName string
	}{
		svInfo.CommitHash, svInfo.CurrentVersion, pn, vn,
	}

	if e := tmpl.ExecuteTemplate(svData, "sv", placeholders); e != nil {
		return nil, e
	}

	return svData.Buf, nil
}

// saveSemverInfo Write the info to a file.
func saveSemverInfo(filename string, info []byte) error {
	if e := os.WriteFile(filename, info, dirMode); e != nil {
		return fmt.Errorf("could not write file %q, reason: %v", filename, e.Error())
	}
	return nil
}

// getCurrentVersion list all versions, sort by semantic version, then give you the one off the top.
func getCurrentVersion(repoPath string) (latestVersion string) {
	latestVersion = "HEAD"

	// Default to HEAD when no tag.
	sco, sce, exitCode, err3 := help.RunRepoCmd(repoPath, "tag", "--sort=-version:refname")
	if err3 != nil || sce != nil || exitCode != 0 {
		latestVersion = ""
		return
	}

	versionsData := bytes.Trim(sco, "\n")
	if versionsData != nil && len(versionsData) > 0 {
		// Split output into an array by newline.
		versions := bytes.Split(versionsData, []byte("\n"))
		latestVersion = string(versions[0])
	}

	return
}

// getCommitHash returns the git commit has for the given tag.
func getCommitHash(repoPath, tag string) (commitHash string, err error) {
	sco, sce, exitCode, err1 := help.RunRepoCmd(repoPath, "rev-list", "-n", "1", tag)
	if err1 != nil {
		return
	}

	if sce == nil && exitCode == 0 {
		commitHashData := bytes.Trim(sco, "\n")
		if len(commitHashData) > 0 {
			commitHash = string(commitHashData)
		}
	}
	return
}

// getNextVersion Get the next calculated semantic version.
func getNextVersion(repoPath, tag string) (nextVer, nextVerReason string) {
	nextVer = "0.1.0"
	revRange := ""
	// Reduce the revision range from the last tag to the most recent commit.
	if tag != "HEAD" {
		revRange = tag + "..HEAD"
	}

	var sco []byte
	var sce, err1 error
	var exitCode int
	// TODO: Handle really large amounts of git logs efficiently.
	// Look at all commit logs since the last tag (maybe only annotated):
	if revRange == "" {
		sco, sce, exitCode, err1 = help.RunRepoCmd(repoPath, "log", "--format=medium")
	} else {
		sco, sce, exitCode, err1 = help.RunRepoCmd(repoPath, "log", "--format=medium", revRange)
		if err1 != nil {
			_ = fmt.Errorf("error retrieving git logs: %v", err1.Error())
			return
		}
	}

	if sce != nil && exitCode != 0 {
		return
	}

	commitLogs := string(sco)
	if len(commitLogs) < 0 {
		return
	}

	re := regexp.MustCompile(`rel:\s*(\d+\.\d+\.\d+)(-.+)?`)
	res := re.FindStringSubmatch(commitLogs)
	// Look for commit message format "rel: x.x.x-optional"
	if len(res) > 0 {
		nextVerReason = "`rel:` type was found in the git logs from the last release to the current HEAD"
		nextVer = res[1] + res[2]
		return
	}

	ver := strings.Split(tag, ".")
	if tag == "HEAD" || len(ver) < 3 {
		nextVerReason = "no previous tags detected"
		return
	}

	// Look for "BREAKING CHANGE" to increment major number
	if strings.Contains(commitLogs, "BREAKING CHANGE\n") {
		ver[0], nextVer, nextVerReason = incrementNumber(ver[0], nextVer, "`BREAKING CHANGE` keyword found in git logs")
		ver[1] = "0"
		ver[2] = "0"

	} else if strings.Contains(commitLogs, "add: ") { // Look for "add:" to increment the minor number
		ver[1], nextVer, nextVerReason = incrementNumber(ver[1], nextVer, "add: keyword found in git logs")
		ver[2] = "0"

	} else { // Increment patch number
		ver[2], nextVer, nextVerReason = incrementNumber(ver[2], nextVer, "no new features or breaking change detected in the git logs")
	}
	nextVer = strings.Join(ver, ".")

	return
}

// incrementNumber add 1 to a numeric string, on failure return numeric number and the reason it failed.
func incrementNumber(a, nv, nvr string) (string, string, string) {
	ret, err1 := strconv.ParseInt(a, 10, 64)
	if err1 != nil {
		return a, "", fmt.Sprintf("unable to increment %q numeric string by 1, reason: %v", a, err1.Error())
	}

	return strconv.FormatInt(ret+1, 10), nv, nvr
}

// incrementNumber add 1 to a numeric string, on failure return numeric number and the reason it failed.
func scrubNumber(a string) string {
	re := regexp.MustCompile(`^(\d+)(-.+)?`)
	res := re.FindStringSubmatch(a)
	dbugf("\na = %v and res = %v\n", a, res)
	if len(res) > 0 {
		dbugf("\nfound %q in %q of length %v\n", res[1], res[0], len(res))
		//fmt.Printf("\nfound %q in %q of length %v\n", res[1], res[0], len(res))
		return res[1]
	}

	return a
}
