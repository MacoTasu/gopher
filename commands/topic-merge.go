package commands

import (
	"../config"
	"../git"
	"fmt"
)

type TopicMergeOpts struct {
	BranchName string
	Config     config.ConfData
}

func TopicMerge(args []string, conf config.ConfData) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("not enough argument")
	}

	tm := &TopicMergeOpts{BranchName: args[0], Config: conf}
	return tm.Exec()
}

// TODO : issue検索してぶら下がってるPRをmergeするのがよい
func (tm *TopicMergeOpts) Exec() (string, error) {
	git := &git.Git{WorkDir: tm.Config.GitWorkDir}

	if _, err := git.Fetch(); err != nil {
		return "", err
	}

	baseBranch := "topic/" + tm.BranchName

	if _, err := git.CheckoutBranch(baseBranch); err != nil {
		return "", err
	}

	if _, err := git.Pull(); err != nil {
		return "", err
	}

	if _, err := git.Merge("origin/" + baseBranch + "-masterdata"); err != nil {
		return "", err
	}

	if _, err := git.Merge("origin/" + baseBranch + "-assetbundle"); err != nil {
		return "", err
	}

	git.Push()

	if _, err := git.CheckoutBranch("master"); err != nil {
		return "", err
	}

	if _, err := git.Pull(); err != nil {
		return "", err
	}

	if _, err := git.Merge(baseBranch); err != nil {
		return "", err
	}

	git.Push()

	return tm.BranchName + " をmergeしたよ", nil
}
