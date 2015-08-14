package mirage

import (
	"../cmd"
	"../config"
	"../util"
	"fmt"
)

type Mirage struct {
	Subdomain  string
	BranchName string
}

func (m *Mirage) Launch() (string, error) {
	percentage, err := util.FreeMemoryPercentage()
	if err != nil {
		return "", err
	}

	if percentage <= 25 {
		return "", fmt.Errorf(fmt.Sprintf("Can't launch. AvailableMemory: %d \\%", percentage))
	}

	conf := config.LoadConfig()

	c := cmd.Cmd{
		Name: "curl",
		Args: []string{conf.MirageUrl, "-d", "subdomain=" + m.Subdomain, "-d", "branch=" + m.BranchName, "-d", "image=" + conf.DockerImage},
	}

	return c.Exec()
}
