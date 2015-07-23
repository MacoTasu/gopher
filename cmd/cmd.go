package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Cmd struct {
	Name string
	Args []string
}

func (c *Cmd) Exec() (string, error) {
	out, err := exec.Command(c.Name, c.Args...).CombinedOutput()

	if err != nil {
		log.Print(string(out))
		return "", fmt.Errorf("%v : %s", err, string(out))
	}

	return strings.TrimSpace(string(out)), nil
}
