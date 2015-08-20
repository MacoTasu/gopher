package commands

import (
	"../config"
	"../git"
	"fmt"
	"strings"
)

type YoshinaOpts struct {
	Subdomain string
	Branches  []string
}

func Yoshina(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("not enough argument")
	}

	branches := strings.Split(args[1], ",")
	if len(branches) < 1 {
		return "", fmt.Errorf("please select branches")
	}

	y := &YoshinaOpts{Subdomain: args[0], Branches: branches}
	return y.Exec()
}

func (y *YoshinaOpts) Exec() (string, error) {
	conf := config.LoadConfig()
	git := &git.Git{WorkDir: conf.GitWorkDir}

	if _, err := git.Fetch(); err != nil {
		return "", err
	}

	baseBranch := y.Branches[0]
	if _, err := git.CheckoutBranch(baseBranch); err != nil {
		return "", err
	}

	if _, err := git.Pull(); err != nil {
		return "", err
	}

	for index, branch := range y.Branches {
		if index == 0 {
			continue
		}

		if _, err := git.Merge("origin/" + branch); err != nil {
			return "", err
		}
	}

	git.PushRemote(baseBranch)

	message := fmt.Sprintf("akane: launch %s %s", y.Subdomain, baseBranch)
	return message, nil
}
