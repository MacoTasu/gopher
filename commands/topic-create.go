package commands

import (
	"../config"
	"../git"
	"code.google.com/p/goauth2/oauth"
	"fmt"
	"github.com/google/go-github/github"
	"log"
)

type TopicCreateOpts struct {
	Prefix      string
	BranchName  string
	IssueNumber string
}

type TopicCreateRequest struct {
	Owner string
	Repo  string
	Base  string
	Head  string
	Title string
	Body  string
}

func TopicCreate(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("not enough argument")
	}

	tc := &TopicCreateOpts{Prefix: "topic/", BranchName: args[0], IssueNumber: args[1]}
	return tc.Exec()
}

func (tc *TopicCreateOpts) Exec() (string, error) {
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

	if _, err := git.Fetch(); err != nil {
		return "", err
	}

	if _, err := git.CheckoutBranch("master"); err != nil {
		return "", err
	}

	if _, err := git.Pull(); err != nil {
		return "", err
	}

	labels := conf.PullRequestLabels

	if err := tc.createAndPullRequest(git, client, &TopicCreateRequest{
		Owner: owner,
		Repo:  repo,
		Base:  "master",
		Head:  tc.baseBranchName(),
		Title: tc.baseBranchName(),
		Body:  conf.PullRequestComment + tc.IssueNumber,
	}, labels); err != nil {
		return "", err
	}

	if err := tc.createAndPullRequest(git, client, &TopicCreateRequest{
		Owner: owner,
		Repo:  repo,
		Base:  tc.baseBranchName(),
		Head:  tc.baseBranchName() + "-masterdata",
		Title: tc.baseBranchName() + "-masterdata",
		Body:  "",
	}, labels); err != nil {
		return "", err
	}

	if _, err := git.CheckoutBranch(tc.baseBranchName()); err != nil {
		return "", err
	}

	if err := tc.createAndPullRequest(git, client, &TopicCreateRequest{
		Owner: owner,
		Repo:  repo,
		Base:  tc.baseBranchName(),
		Head:  tc.baseBranchName() + "-assetbundle",
		Title: tc.baseBranchName() + "-assetbundle",
		Body:  "",
	}, labels); err != nil {
		return "", err
	}

	return "ʕ ◔ϖ◔ʔ < ブランチ作成したよ", nil
}

func (tc *TopicCreateOpts) createAndPullRequest(git *git.Git, client *github.Client, topicCreateRequest *TopicCreateRequest, labels []string) (err error) {
	branchName := topicCreateRequest.Head
	git.CreateBranch(branchName)
	git.EmptyCommit()
	git.PushRemote(branchName)

	owner := topicCreateRequest.Owner
	repo := topicCreateRequest.Repo

	head := owner + ":" + branchName
	base := topicCreateRequest.Base
	title := topicCreateRequest.Title
	body := topicCreateRequest.Body
	pull, _, _ := client.PullRequests.Create(owner, repo, &github.NewPullRequest{
		Title: &title,
		Head:  &head,
		Base:  &base,
		Body:  &body,
	})

	_, _, label_err := client.Issues.AddLabelsToIssue(owner, repo, *pull.Number, labels)
	if label_err != nil {
		log.Println(label_err)
	}

	return nil
}

func (tc *TopicCreateOpts) baseBranchName() string {
	return tc.Prefix + tc.BranchName
}
