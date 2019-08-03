package github

import (
	"context"
	"regexp"
	"strings"

	"github.com/MacoTasu/gopher/config"
	"github.com/MacoTasu/gopher/git"
	gh "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Github struct {
	Client *gh.Client
}

func New(ctx context.Context, config config.ConfData) (*Github, error) {
	git := &git.Git{WorkDir: config.GitWorkDir}

	token, err := git.FetchAccessToken("gopher.token")
	if err != nil {
		return nil, err
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := gh.NewClient(tc)

	return &Github{Client: client}, nil
}

func (github *Github) FetchPullRequestHeadRef(IssueNumber int, owner string, repo string) ([]string, error) {
	ctx := context.Background()
	pull, _, err := github.Client.PullRequests.Get(ctx, owner, repo, IssueNumber)
	if err != nil {
		return nil, err
	}

	pattern := `merge: ([a-zA-Z0-9,/\-_\.]+)`
	result := regexp.MustCompile(pattern).FindString(*pull.Body)
	mergeBranchName := strings.Replace(result, "merge: ", "", -1)
	branches := strings.Split(mergeBranchName, ",")

	// mergeが見つからなかった時は、refを使う
	if branches[0] == "" {
		ref := *pull.Head.Ref
		branches[0] = ref
		branches = append(branches, ref+"-masterdata")
		branches = append(branches, ref+"-assetbundle")
	}

	return branches, nil
}
