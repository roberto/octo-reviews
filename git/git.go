package git

import (
	"bytes"
	"os/exec"
	"regexp"
	"strings"
)

type LocalInfo struct {
	Branch  string
	Owner   string
	Project string
}

func NewLocalInfo() LocalInfo {
	branch := getBranchName()
	owner, project := getOwnerAndProject()

	return LocalInfo{branch, owner, project}
}

func getBranchName() string {
	return executeGitCommand("rev-parse", "--abbrev-ref", "HEAD")
}

func getOwnerAndProject() (string, string) {
	repositoryUrl := getRepositoryUrl()
	re := regexp.MustCompile(`:(.*)\/(.*)\.`)
	matches := re.FindAllStringSubmatch(repositoryUrl, -1)[0]
	return matches[1], matches[2]
}

func getRepositoryUrl() string {
	return executeGitCommand("config", "--get", "remote.origin.url")
}

func executeGitCommand(args ...string) string {
	gitPath, _ := exec.LookPath("git")
	cmd := exec.Command(gitPath, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	return strings.TrimSpace(out.String())
}
