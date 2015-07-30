package commands

import (
	"../config"
	"../models"
	"code.google.com/p/goauth2/oauth"
	"fmt"
	"github.com/google/go-github/github"
	"strconv"
	"strings"
	"time"
)

type TopicDeploy struct {
	ServerName  string
	IssueNumber int
	Options     []string
}

func topicDeploy(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("not enough argument")
	}

	number, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}

	td := &TopicDeploy{ServerName: args[0], IssueNumber: number, Options: args[2:]}
	return td.Exec()
}

func (td *TopicDeploy) Exec() (string, error) {
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

	ref, err := td.fetchPullRequestHeadRef(client, owner, repo)
	if err != nil {
		return "", err
	}

	now := time.Now()
	git.Fetch()
	git.CheckoutBranch(ref)
	deployRefName := fmt.Sprintf("jenkins/%s-%02d%02d.%02d%02d", ref, now.Month(), now.Day(), now.Hour(), now.Minute())
	git.CreateBranch(deployRefName)
	git.Merge("origin/" + ref + "-masterdata")
	git.Merge("origin/" + ref + "-assetbundle")
	git.PushRemote(deployRefName)

	message := fmt.Sprintf("akane: deploy %s %s", td.ServerName, deployRefName)
	if len(td.Options) > 0 {
		message = fmt.Sprintf("%s %s", message, strings.Join(td.Options, " "))
	}

	return message, nil
}

func (td *TopicDeploy) fetchPullRequestHeadRef(client *github.Client, owner string, repo string) (string, error) {
	pull, _, err := client.PullRequests.Get(owner, repo, td.IssueNumber)
	if err != nil {
		return "", err
	}

	return *pull.Head.Ref, nil
}
