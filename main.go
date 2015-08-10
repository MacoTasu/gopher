package main

import (
	"./commands"
	"./config"
	"fmt"
	"github.com/m0t0k1ch1/ape"
	"log"
)

func main() {
	conf := config.LoadConfig()
	con := ape.NewConnection("gopher", "gopher")
	con.UseTLS = true
	con.Password = conf.Password
	prefix := "ʕ ◔ϖ◔ʔ "

	if err := con.Connect(conf.Server); err != nil {
		log.Fatal(err)
	}

	con.RegisterChannel(conf.Channel)

	con.AddAction("topic-deploy", func(e *ape.Event) {
		con.SendMessage(prefix + "<" + "topic-deploy not ybsk")
		result, err := commands.TopicDeploy(e.Command().Args())
		if err != nil {
			con.SendMessage(prefix + "< " + fmt.Sprintln(err))
		} else {
			con.SendMessage(result)
		}
	})

	con.AddAction("topic-create", func(e *ape.Event) {
		con.SendMessage(prefix + "<" + "topic-create not ybsk")
		result, err := commands.TopicCreate(e.Command().Args())
		if err != nil {
			con.SendMessage(prefix + "< " + fmt.Sprintln(err))
		} else {
			con.SendMessage(result)
		}
	})

	con.AddAction("topic-launch", func(e *ape.Event) {
		con.SendMessage(prefix + "<" + "topic-launch not ybsk")
		result, err := commands.TopicLaunch(e.Command().Args())
		if err != nil {
			con.SendMessage(prefix + "< " + fmt.Sprintln(err))
		} else {
			con.SendMessage(result)
		}
	})

	con.AddAction("launch", func(e *ape.Event) {
		con.SendMessage(prefix + "<" + "launch not ybsk")
		result, err := commands.Launch(e.Command().Args())
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

	con.AddDefaultAction(func(e *ape.Event) {
		con.SendMessage(prefix + "？？")
	})

	con.AddAction("ほんわかぽにゃぽにゃ", func(e *ape.Event) {
		con.Part(con.Channel())
		con.Join(con.Channel())
	})

	con.Loop()
}
