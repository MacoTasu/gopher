package commands

import (
	"../cmd"
	"../config"
	"../models"
	"code.google.com/p/goauth2/oauth"
	"fmt"
	"github.com/google/go-github/github"
	"strconv"
	"time"
)

type TopicLaunch struct {
	Subdomain   string
	IssueNumber int
}

func topicLaunch(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("not enough argument")
	}

	number, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}

	tl := &TopicLaunch{Subdomain: args[0], IssueNumber: number}
	return tl.Exec()
}

func (tl *TopicLaunch) Exec() (string, error) {
	conf := config.LoadConfig()
	git := &models.Git{WorkDir: conf.GitWorkDir}

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

	now := time.Now()
	git.Fetch()
	git.CheckoutBranch(ref)
	deployRefName := fmt.Sprintf("%s-%02d%02d.%02d%02d", ref, now.Month(), now.Day(), now.Hour(), now.Minute())
	git.CreateBranch(deployRefName)
	git.Merge("origin/" + ref + "-masterdata")
	git.Merge("origin/" + ref + "-assetbundle")
	git.PushRemote(deployRefName)

	c := cmd.Cmd{
		Name: "curl",
		Args: []string{conf.MirageUrl, "-d", "subdomain=" + tl.Subdomain, "-d", "branch=" + deployRefName, "-d", "image=" + conf.DockerImage},
	}

	if _, err := c.Exec(); err != nil {
		return "", err
	}

	return fmt.Sprintf("ʕ ◔ϖ◔ʔ < %sに%sのブランチで環境作成したよ！", tl.Subdomain, deployRefName), nil
}

func (tl *TopicLaunch) fetchPullRequestHeadRef(client *github.Client, owner string, repo string) (string, error) {
	pull, _, err := client.PullRequests.Get(owner, repo, tl.IssueNumber)
	if err != nil {
		return "", err
	}

	return *pull.Head.Ref, nil
}
