package main

import (
	"fmt"

	"./commands"
	"./config"
	"github.com/shogo82148/ape-slack"
)

func main() {
	conf := config.LoadConfig()
	con := ape.NewConnection(conf.Token)
	prefix := "ʕ ◔ϖ◔ʔ "

	con.RegisterChannel(conf.Channel)

	con.AddAction("topic-merge", func(e *ape.Event) {
		con.SendMessage(prefix + "< " + "topic-merge not ybsk")
		result, err := commands.TopicMerge(e.Command().Args())
		if err != nil {
			con.SendMessage(prefix + "< " + fmt.Sprintln(err))
		} else {
			con.SendMessage(result)
		}
	})

	con.AddAction("topic-deploy", func(e *ape.Event) {
		con.SendMessage(prefix + "< " + "topic-deploy not ybsk")
		result, err := commands.TopicDeploy(e.Command().Args())
		if err != nil {
			con.SendMessage(prefix + "< " + fmt.Sprintln(err))
		} else {
			con.SendMessage(result)
		}
	})

	con.AddAction("topic-create", func(e *ape.Event) {
		con.SendMessage(prefix + "< " + "topic-create not ybsk")
		result, err := commands.TopicCreate(e.Command().Args())
		if err != nil {
			con.SendMessage(prefix + "< " + fmt.Sprintln(err))
		} else {
			con.SendMessage(result)
		}
	})

	con.AddAction("topic-launch", func(e *ape.Event) {
		con.SendMessage(prefix + "< " + "topic-launch not ybsk")
		result, err := commands.TopicLaunch(e.Command().Args())
		if err != nil {
			con.SendMessage(prefix + "< " + fmt.Sprintln(err))
		} else {
			con.SendMessage(result)
		}
	})

	con.AddAction("launch", func(e *ape.Event) {
		con.SendMessage(prefix + "< " + "launch not ybsk")
		result, err := commands.Launch(e.Command().Args())
		if err != nil {
			con.SendMessage(prefix + "< " + fmt.Sprintln(err))
		} else {
			con.SendMessage(result)
		}
	})

	con.AddAction("deploy", func(e *ape.Event) {
		con.SendMessage(prefix + "< " + "deploy not ybsk")
		result, err := commands.Deploy(e.Command().Args())
		if err != nil {
			con.SendMessage(prefix + "< " + fmt.Sprintln(err))
		} else {
			con.SendMessage(result)
		}
	})

	con.AddAction("help", func(e *ape.Event) {
		result, err := commands.Help(e.Command().Args())
		if err != nil {
			con.SendMessage(prefix + "< " + fmt.Sprintln(err))
		} else {
			con.SendMessage(result)
		}
	})

	con.AddAction("pray", func(e *ape.Event) {
		con.SendMessage(prefix + "< きっと大丈夫やで")
	})

	con.AddAction("すき", func(e *ape.Event) {
		con.SendMessage("https://qiita-image-store.s3.amazonaws.com/0/14952/b7acefb3-d354-ea13-2242-f2c16919793d.png")
	})

	con.Loop()
}
