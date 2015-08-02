package cmd

import (
	"github.com/github/hub/Godeps/_workspace/src/github.com/bmizerany/assert"

	"../cmd"
	"testing"
)

func TestExec(t *testing.T) {
	execCmd := &cmd.Cmd{Name: "echo", Args: []string{"hoge"}}
	result, _ := execCmd.Exec()
	assert.Equal(t, "hoge", result)
}
