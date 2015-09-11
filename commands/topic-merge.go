package commands

import (
	"../config"
	"../git"
	"fmt"
)

type TopicMergeOpts struct {
	BranchName string
}

func TopicMerge(args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("not enough argument")
	}

	tm := &TopicMergeOpts{BranchName: args[0]}
	return tm.Exec()
}

// TODO : issue検索してぶら下がってるPRをmergeするのがよい
func (tm *TopicMergeOpts) Exec() (string, error) {
	conf := config.LoadConfig()
	git := &git.Git{WorkDir: conf.GitWorkDir}

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

	git.PushRemote(baseBranch)

	if _, err := git.CheckoutBranch("master"); err != nil {
		return "", err
	}

	if _, err := git.Pull(); err != nil {
		return "", err
	}

	if _, err := git.Merge(baseBranch); err != nil {
		return "", err
	}

	return tm.BranchName + " をmergeしたよ", nil
}
