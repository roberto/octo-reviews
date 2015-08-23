package git

import (
	"bytes"
	"os/exec"
	"strings"
)

func executeGitCommand(args ...string) string {
	gitPath, _ := exec.LookPath("git")
	cmd := exec.Command(gitPath, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	return strings.TrimSpace(out.String())
}

func BranchName() string {
	return executeGitCommand("rev-parse", "--abbrev-ref", "HEAD")
}

func RepositoryUrl() string {
	return executeGitCommand("config", "--get", "remote.origin.url")
}
