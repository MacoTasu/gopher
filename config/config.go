package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ConfData struct {
	Token              string   `yaml:"token"`
	Channel            string   `yaml:"channel"`
	GitWorkDir         string   `yaml:"git_work_dir"`
	PullRequestComment string   `yaml:"pull_request_comment"`
	MirageUrl          string   `yaml:"mirage_url"`
	DockerImage        string   `yaml:"docker_image"`
	PullRequestLabels  []string `yaml:"pull_request_labels"`
}

// conf.ymlをロードし構造体へ格納
func LoadConfig(path string) *ConfData {
	buf, err := ioutil.ReadFile(path)

	d := ConfData{}
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(buf, &d); err != nil {
		panic(err)
	}

	return &d
}
