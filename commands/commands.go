package commands

import (
	"fmt"
)

type Command struct {
	Run  func([]string) (string, error)
	Args []string
}

func (c *Command) FetchFunc(args []string) error {
	err := c.fetchSubCommand(args)
	if err != nil {
		return err
	}

	return nil
}

func (c *Command) fetchSubCommand(args []string) error {
	if len(args) <= 0 {
		return fmt.Errorf("please choose the command")
	}

	//TODO : 動的につくったほうがいい
	subCommands := map[string]func([]string) (string, error){
		"topic-create": topicCreate,
		"topic-deploy": topicDeploy,
		"topic-launch": topicLaunch,
		"help":         help,
		"launch":       launch,
	}

	subCommand := subCommands[args[0]]

	if subCommand == nil {
		return fmt.Errorf("unknown command %s", args[0])
	}

	c.Run = subCommand
	c.Args = args[1:]

	return nil
}

func (c *Command) Call() (string, error) {
	return c.Run(c.Args)
}
