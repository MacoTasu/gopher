package github

import (
	"../config"
	"../git"
	"code.google.com/p/goauth2/oauth"
	gh "github.com/google/go-github/github"
	"regexp"
	"strings"
)

type Github struct {
	Client *gh.Client
}

func New(config config.ConfData) (*Github, error) {
	git := &git.Git{WorkDir: config.GitWorkDir}

	token, err := git.FetchAccessToken("gopher.token")
	if err != nil {
		return nil, err
	}

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}
	client := gh.NewClient(t.Client())
	return &Github{Client: client}, nil
}

func (github *Github) FetchPullRequestHeadRef(IssueNumber int, owner string, repo string) ([]string, error) {
	pull, _, err := github.Client.PullRequests.Get(owner, repo, IssueNumber)
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
