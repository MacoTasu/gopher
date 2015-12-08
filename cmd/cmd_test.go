package cmd

import (
	"testing"

	"../cmd"
	"github.com/github/hub/Godeps/_workspace/src/github.com/bmizerany/assert"
)

func TestExec(t *testing.T) {
	execCmd := &cmd.Cmd{Name: "echo", Args: []string{"hoge"}}
	result, _ := execCmd.Exec()
	assert.Equal(t, "hoge", result)
}
