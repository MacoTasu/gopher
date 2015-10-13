package main

import (
	"fmt"

	"./commands"
	"./config"
	"flag"
	"github.com/shogo82148/ape-slack"
)

var (
	confFile = flag.String("conf", "config.yml", "config file path")
)

func main() {
	flag.Parse()
	conf := config.LoadConfig(*confFile)
	con := ape.NewConnection(conf.Token)
	prefix := "ʕ ◔ϖ◔ʔ "

	con.RegisterChannel(conf.Channel)

	con.AddAction("topic-merge", func(e *ape.Event) {
		con.SendMessage(prefix + "< " + "topic-merge not ybsk")
		result, err := commands.TopicMerge(e.Command().Args(), *conf)
		if err != nil {
			con.SendMessage(prefix + "< " + fmt.Sprintln(err))
		} else {
			con.SendMessage(result)
		}
	})

	con.AddAction("topic-deploy", func(e *ape.Event) {
		con.SendMessage(prefix + "< " + "topic-deploy not ybsk")
		result, err := commands.TopicDeploy(e.Command().Args(), *conf)
		if err != nil {
			con.SendMessage(prefix + "< " + fmt.Sprintln(err))
		} else {
			con.SendMessage(result)
		}
	})

	con.AddAction("topic-create", func(e *ape.Event) {
		con.SendMessage(prefix + "< " + "topic-create not ybsk")
		result, err := commands.TopicCreate(e.Command().Args(), *conf)
		if err != nil {
			con.SendMessage(prefix + "< " + fmt.Sprintln(err))
		} else {
			con.SendMessage(result)
		}
	})

	con.AddAction("topic-launch", func(e *ape.Event) {
		con.SendMessage(prefix + "< " + "topic-launch not ybsk")
		result, err := commands.TopicLaunch(e.Command().Args(), *conf)
		if err != nil {
			con.SendMessage(prefix + "< " + fmt.Sprintln(err))
		} else {
			con.SendMessage(result)
		}
	})

	con.AddAction("launch", func(e *ape.Event) {
		con.SendMessage(prefix + "< " + "launch not ybsk")
		result, err := commands.Launch(e.Command().Args(), *conf)
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

	con.Loop()
}
