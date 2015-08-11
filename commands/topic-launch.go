package commands

import (
	"../config"
	"../git"
	"../mirage"
	"code.google.com/p/goauth2/oauth"
	"fmt"
	"github.com/google/go-github/github"
	"strconv"
	"time"
)

type TopicLaunchOpts struct {
	Subdomain   string
	IssueNumber int
}

func TopicLaunch(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("not enough argument")
	}

	number, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}

	tl := &TopicLaunchOpts{Subdomain: args[0], IssueNumber: number}
	return tl.Exec()
}

func (tl *TopicLaunchOpts) Exec() (string, error) {
	conf := config.LoadConfig()
	git := &git.Git{WorkDir: conf.GitWorkDir}

	token, err := git.FetchAccessToken("gopher.token")
	if err != nil {
		return "", err
	}
	owner, repo, err := git.FetchOwnerAndRepo()
	if err != nil {
		return "", err
	}

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}
	client := github.NewClient(t.Client())

	ref, err := tl.fetchPullRequestHeadRef(client, owner, repo)
	if err != nil {
		return "", err
	}

	if _, err := git.Fetch(); err != nil {
		return "", err
	}

	if _, err := git.CheckoutBranch(ref); err != nil {
		return "", err
	}

	if _, err := git.Pull(); err != nil {
		return "", err
	}

	now := time.Now()
	deployRefName := fmt.Sprintf("gopher/%s-%02d%02d.%02d%02d", ref, now.Month(), now.Day(), now.Hour(), now.Minute())

	if _, err := git.CreateBranch(deployRefName); err != nil {
		return "", err
	}

	if _, err := git.Merge("origin/" + ref + "-masterdata"); err != nil {
		return "", err
	}

	if _, err := git.Merge("origin/" + ref + "-assetbundle"); err != nil {
		return "", err
	}

	git.PushRemote(deployRefName)

	mirage := &mirage.Mirage{Subdomain: tl.Subdomain, BranchName: deployRefName}
	if _, err := mirage.Launch(); err != nil {
		return "", err
	}

	return fmt.Sprintf("ʕ ◔ϖ◔ʔ < %s に %s で環境作成依頼をだしたよ！", tl.Subdomain, deployRefName), nil
}

func (tl *TopicLaunchOpts) fetchPullRequestHeadRef(client *github.Client, owner string, repo string) (string, error) {
	pull, _, err := client.PullRequests.Get(owner, repo, tl.IssueNumber)
	if err != nil {
		return "", err
	}

	return *pull.Head.Ref, nil
}
