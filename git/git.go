package git

import (
	"../cmd"
	"fmt"
	"regexp"
)

type Git struct {
	WorkDir string
}

var (
	re = regexp.MustCompile(`^(?:git@github\.com:|https://github\.com/)([^/]+)/([^/]+?)(?:\.git)?$`)
)

func (g *Git) Fetch() (string, error) {

	c := cmd.Cmd{
		Name: "git",
		Args: g.appendGitOptions([]string{"fetch"}),
	}

	return c.Exec()
}

func (g *Git) Pull() (string, error) {

	c := cmd.Cmd{
		Name: "git",
		Args: g.appendGitOptions([]string{"pull"}),
	}

	return c.Exec()
}

func (g *Git) EmptyCommit() (string, error) {

	c := cmd.Cmd{
		Name: "git",
		Args: g.appendGitOptions([]string{"commit", "--allow-empty", "-m", "empty-commit"}),
	}

	return c.Exec()
}

func (g *Git) CheckoutBranch(branchName string) (string, error) {

	c := cmd.Cmd{
		Name: "git",
		Args: g.appendGitOptions([]string{"checkout", branchName}),
	}

	return c.Exec()
}

func (g *Git) CreateBranch(branchName string) (string, error) {

	c := cmd.Cmd{
		Name: "git",
		Args: g.appendGitOptions([]string{"checkout", "-b", branchName}),
	}

	return c.Exec()
}

func (g *Git) Merge(branchName string) (string, error) {

	c := cmd.Cmd{
		Name: "git",
		Args: g.appendGitOptions([]string{"merge", "--no-ff", branchName}),
	}

	return c.Exec()
}

func (g *Git) PushRemote(branchName string) (string, error) {

	c := cmd.Cmd{
		Name: "git",
		Args: g.appendGitOptions([]string{"push", "-u", "origin", branchName}),
	}

	return c.Exec()
}

func (g *Git) Push() (string, error) {

	c := cmd.Cmd{
		Name: "git",
		Args: g.appendGitOptions([]string{"push"}),
	}

	return c.Exec()
}

func (g *Git) FetchAccessToken(keyName string) (string, error) {

	c := cmd.Cmd{
		Name: "git",
		Args: g.appendGitOptions([]string{"config", "--get", keyName}),
	}

	return c.Exec()
}

// m0t0k1ch1++
func (g *Git) FetchOwnerAndRepo() (string, string, error) {

	c := cmd.Cmd{
		Name: "git",
		Args: g.appendGitOptions([]string{"config", "--get", "remote.origin.url"}),
	}

	url, err := c.Exec()
	if err != nil {
		return "", "", err
	}

	matches := re.FindStringSubmatch(url)
	if len(matches) != 3 {
		err = fmt.Errorf("can't parse remote.origin.url")
		return "", "", err
	}

	return matches[1], matches[2], nil
}

func (g *Git) appendGitOptions(args []string) []string {
	gitOptions := []string{"--git-dir", g.WorkDir + "/.git", "--work-tree", g.WorkDir}
	return append(gitOptions, args...)
}
