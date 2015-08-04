package commands

import (
	"../config"
	"../git"
	"../mirage"
	"fmt"
)

type Launch struct {
	Subdomain  string
	BranchName string
}

func launch(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("not enough argument")
	}

	l := &Launch{Subdomain: args[0], BranchName: args[1]}
	return l.Exec()
}

func (l *Launch) Exec() (string, error) {
	conf := config.LoadConfig()
	git := &git.Git{WorkDir: conf.GitWorkDir}

	if _, err := git.Fetch(); err != nil {
		return "", err
	}

	if _, err := git.CheckoutBranch(l.BranchName); err != nil {
		return "", err
	}

	mirage := &mirage.Mirage{Subdomain: l.Subdomain, BranchName: l.BranchName}
	if _, err := mirage.Launch(); err != nil {
		return "", err
	}

	return fmt.Sprintf("ʕ ◔ϖ◔ʔ < %s に %s で環境作成したよ！", l.Subdomain, l.BranchName), nil
}
