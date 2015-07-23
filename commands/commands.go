package commands

import (
	"fmt"
)

type Command struct {
	Run func([]string) (string, error)
}

func (c *Command) Call(args []string) (string, error) {
	cmd, cmdArgs, err := c.fetchSubCommand(args)
	if err != nil {
		return "", err
	}

	return cmd.Run(cmdArgs)
}

func (c *Command) fetchSubCommand(args []string) (*Command, []string, error) {
	if len(args) <= 0 {
		return nil, nil, fmt.Errorf("please choose the command")
	}

	//TODO : 動的につくったほうがいい
	subCommands := map[string]func([]string) (string, error){
		"topic-create": topicCreate,
		"topic-deploy": topicDeploy,
	}

	subCommand := subCommands[args[0]]

	if subCommand == nil {
		return nil, nil, fmt.Errorf("unknown command %s", args[0])
	}

	return &Command{Run: subCommand}, args[1:], nil
}
