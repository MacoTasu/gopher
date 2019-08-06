package mirage

import (
	"errors"
	"fmt"

	"github.com/MacoTasu/gopher/cmd"
	"github.com/MacoTasu/gopher/util"
)

type Mirage struct {
	Subdomain   string
	BranchName  string
	Url         string
	DockerImage string
}

func (m *Mirage) Launch() (string, error) {
	percentage, err := util.FreeMemoryPercentage()
	if err != nil {
		return "", err
	}

	if percentage <= 15 {
		return "", errors.New(fmt.Sprintf("Can't launch. AvailableMemory: %d%%\n", percentage))
	}

	c := cmd.Cmd{
		Name: "curl",
		Args: []string{m.Url, "-d", "subdomain=" + m.Subdomain, "-d", "branch=" + m.BranchName, "-d", "image=" + m.DockerImage},
	}

	return c.Exec()
}
