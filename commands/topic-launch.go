package commands

import (
	"../config"
	"../git"
	"../github"
	"../mirage"
	"fmt"
	"strconv"
	"time"
)

type TopicLaunchOpts struct {
	Subdomain   string
	IssueNumber int
	Config      config.ConfData
}

func TopicLaunch(args []string, conf config.ConfData) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("not enough argument")
	}

	number, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}

	tl := &TopicLaunchOpts{Subdomain: args[0], IssueNumber: number, Config: conf}
	return tl.Exec()
}

func (tl *TopicLaunchOpts) Exec() (string, error) {
	git := &git.Git{WorkDir: tl.Config.GitWorkDir}

	owner, repo, err := git.FetchOwnerAndRepo()
	if err != nil {
		return "", err
	}

	github, err := github.New(tl.Config)
	if err != nil {
		return "", err
	}

	branches, err := github.FetchPullRequestHeadRef(tl.IssueNumber, owner, repo)
	baseBranch := branches[0]
	if _, err := git.Fetch(); err != nil {
		return "", err
	}

	if _, err := git.CheckoutBranch(baseBranch); err != nil {
		return "", err
	}

	if _, err := git.Pull(); err != nil {
		return "", err
	}

	now := time.Now()
	deployRefName := fmt.Sprintf("gopher/%s-%02d%02d.%02d%02d", baseBranch, now.Month(), now.Day(), now.Hour(), now.Minute())

	if _, err := git.CreateBranch(deployRefName); err != nil {
		return "", err
	}

	for index, branch := range branches {
		if index == 0 {
			continue
		}

		if _, err := git.Merge("origin/" + branch); err != nil {
			return "", err
		}
	}

	git.PushRemote(deployRefName)

	mirage := &mirage.Mirage{Subdomain: tl.Subdomain, BranchName: deployRefName, Url: tl.Config.MirageUrl, DockerImage: tl.Config.DockerImage}
	if _, err := mirage.Launch(); err != nil {
		return "", err
	}

	return fmt.Sprintf("ʕ ◔ϖ◔ʔ < %s に %s で環境作成依頼をだしたよ！", tl.Subdomain, deployRefName), nil
}
