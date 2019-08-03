package cmd

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestExec(t *testing.T) {
	execCmd := &Cmd{Name: "echo", Args: []string{"hoge"}}
	result, _ := execCmd.Exec()
	assert.Equal(t, "hoge", result)
}
