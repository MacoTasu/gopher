package commands

import (
	"testing"

	"../commands"
	"github.com/github/hub/Godeps/_workspace/src/github.com/bmizerany/assert"
)

func TestHelp(t *testing.T) {
	cmd := &commands.Command{}
	cmd.FetchFunc([]string{"help"})
	result, _ := cmd.Call()
	assert.Equal(t, "ʕ ◔ϖ◔ʔ < https://github.com/MacoTasu/gopher#gopher", result)
}
