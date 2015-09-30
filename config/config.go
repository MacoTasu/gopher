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
	AfterImage         string   `yaml:"after_image"`
	PullRequestLabels  []string `yaml:"pull_request_labels"`
}

// conf.ymlをロードし構造体へ格納
func LoadConfig() *ConfData {
	buf, err := ioutil.ReadFile("config/config_local.yml")

	d := ConfData{}
	if err != nil {
		buf, err := ioutil.ReadFile("config/config.yml")
		if err != nil {
			panic(err)
		}

		if err := yaml.Unmarshal(buf, &d); err != nil {
			panic(err)
		}
		return &d
	}

	if err := yaml.Unmarshal(buf, &d); err != nil {
		panic(err)
	}

	return &d
}
