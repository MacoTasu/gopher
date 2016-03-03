package commands

import (
	"../config"
	"../git"
	"../jenkins"
	"../mirage"
	"bytes"
	"fmt"
	"net/url"
	"text/template"
)

type LaunchOpts struct {
	Subdomain  string
	BranchName string
	Config     config.ConfData
	Launcher   string
}

func Launch(args []string, conf config.ConfData, launcher string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("not enough argument")
	}

	l := &LaunchOpts{Subdomain: args[0], BranchName: args[1], Config: conf, Launcher: launcher}
	return l.Exec()
}

func (l *LaunchOpts) Exec() (string, error) {
	git := &git.Git{WorkDir: l.Config.GitWorkDir}

	if _, err := git.Fetch(); err != nil {
		return "", err
	}

	if _, err := git.CheckoutBranch(l.BranchName); err != nil {
		return "", err
	}

	if l.Config.Jenkins.Url != "" {
		return l.execOnJenkins()
	} else {
		return l.execOnMirage()
	}
}

func (l *LaunchOpts) execOnJenkins() (string, error) {
	v := url.Values{}
	for key, value := range l.Config.Jenkins.LaunchParameters {
		tmpl, err := template.New("launch").Parse(value)
		if err != nil {
			return "", fmt.Errorf("jenkins launch parameters template parse error: %s", err)
		}
		b := new(bytes.Buffer)
		err = tmpl.Execute(b, l)
		if err != nil {
			return "", fmt.Errorf("jenkins launch parameters template execute error: %s", err)
		}
		v.Add(key, b.String())
	}
	j := &jenkins.Jenkins{Config: l.Config}
	bp := jenkins.BuildParameters{
		TaskName:   l.Config.Jenkins.LaunchTask,
		Parameters: v,
	}
	if err := j.Build(bp); err != nil {
		return "", err
	}
	return fmt.Sprintf("ʕ ◔ϖ◔ʔ < %s に %s で環境作成依頼をだしたよ！", l.Subdomain, l.BranchName), nil
}

func (l *LaunchOpts) execOnMirage() (string, error) {
	mirage := &mirage.Mirage{Subdomain: l.Subdomain, BranchName: l.BranchName, Url: l.Config.MirageUrl, DockerImage: l.Config.DockerImage}
	if _, err := mirage.Launch(); err != nil {
		return "", err
	}
	return fmt.Sprintf("ʕ ◔ϖ◔ʔ < %s に %s で環境作成依頼をだしたよ！", l.Subdomain, l.BranchName), nil
}
