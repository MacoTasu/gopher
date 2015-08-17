package commands

import (
	"fmt"
	"strings"
)

type DeployOpts struct {
	ServerName string
	BranchName string
	Options    []string
}

func Deploy(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("not enough argument")
	}

	l := &DeployOpts{ServerName: args[0], BranchName: args[1], Options: args[2:]}
	return l.Exec()
}

func (d *DeployOpts) Exec() (string, error) {

	message := fmt.Sprintf("akane: deploy %s %s", d.ServerName, d.BranchName)
	if len(d.Options) > 0 {
		message = fmt.Sprintf("%s %s", message, strings.Join(d.Options, " "))
	}

	return message, nil
}
