package commands

import (
	"../config"
	"../git"
	"../github"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type TopicDeployOpts struct {
	ServerName  string
	IssueNumber int
	Options     []string
	Config      config.ConfData
}

func TopicDeploy(args []string, conf config.ConfData) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("not enough argument")
	}

	number, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}

	td := &TopicDeployOpts{ServerName: args[0], IssueNumber: number, Options: args[2:], Config: conf}
	return td.Exec()
}

func (td *TopicDeployOpts) Exec() (string, error) {
	git := &git.Git{WorkDir: td.Config.GitWorkDir}

	owner, repo, err := git.FetchOwnerAndRepo()
	if err != nil {
		return "", err
	}

	github, err := github.New(td.Config)
	if err != nil {
		return "", err
	}

	branches, err := github.FetchPullRequestHeadRef(td.IssueNumber, owner, repo)
	if err != nil {
		return "", err
	}

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
	deployRefName := fmt.Sprintf("jenkins/%s-%02d%02d.%02d%02d", baseBranch, now.Month(), now.Day(), now.Hour(), now.Minute())

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

	message := fmt.Sprintf("akane: deploy %s %s", td.ServerName, deployRefName)
	if len(td.Options) > 0 {
		message = fmt.Sprintf("%s %s", message, strings.Join(td.Options, " "))
	}

	return message, nil
}
