package commands

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/MacoTasu/gopher/config"
	"github.com/MacoTasu/gopher/git"
	"github.com/MacoTasu/gopher/github"
)

type TopicLaunchOpts struct {
	Subdomain   string
	IssueNumber int
	Config      config.ConfData
	Launcher    string
	ExtraArgs   []string
}

func TopicLaunch(ctx context.Context, args []string, conf config.ConfData, launcher string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("not enough argument")
	}

	number, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}

	tl := &TopicLaunchOpts{
		Subdomain:   args[0],
		IssueNumber: number,
		Config:      conf,
		Launcher:    launcher,
	}
	if len(args) > 2 {
		tl.ExtraArgs = args[2:]
	}
	return tl.Exec(ctx)
}

func (tl *TopicLaunchOpts) Exec(ctx context.Context) (string, error) {
	git := &git.Git{WorkDir: tl.Config.GitWorkDir}
	defer git.Reset("HEAD", true)

	owner, repo, err := git.FetchOwnerAndRepo()
	if err != nil {
		return "", err
	}

	github, err := github.New(ctx, tl.Config)
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

	lo := &LaunchOpts{
		Subdomain:  tl.Subdomain,
		BranchName: deployRefName,
		Config:     tl.Config,
		Launcher:   tl.Launcher,
		ExtraArgs:  tl.ExtraArgs,
	}

	return lo.Exec()
}
