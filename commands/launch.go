package commands

import (
	"fmt"

	"../config"
	"../git"
	"../mirage"
)

type LaunchOpts struct {
	Subdomain  string
	BranchName string
	Config     config.ConfData
}

func Launch(args []string, conf config.ConfData) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("not enough argument")
	}

	l := &LaunchOpts{Subdomain: args[0], BranchName: args[1], Config: conf}
	return l.Exec()
}

func (l *LaunchOpts) Exec() (string, error) {
	git := &git.Git{WorkDir: l.Config.GitWorkDir}

	if _, err := git.Fetch(); err != nil {
		return "", err
	}

	if _, err := git.CheckoutBranch(l.BranchName); err != nil {
		return "", err
	}

	mirage := &mirage.Mirage{Subdomain: l.Subdomain, BranchName: l.BranchName, Url: l.Config.MirageUrl, DockerImage: l.Config.DockerImage}
	if _, err := mirage.Launch(); err != nil {
		return "", err
	}

	return fmt.Sprintf("ʕ ◔ϖ◔ʔ < %s に %s で環境作成依頼をだしたよ！", l.Subdomain, l.BranchName), nil
}
