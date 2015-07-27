package main

import (
	"./commands"
	"./config"
	"fmt"
	"github.com/thoj/go-ircevent"
	"log"
	"strings"
)

func main() {
	conf := config.LoadConfig()
	ircobj := irc.IRC("gopher", "gopher")
	ircobj.UseTLS = true
	ircobj.Password = conf.Password
	prefix := "ʕ ◔ϖ◔ʔ < "
	if err := ircobj.Connect(conf.Server); err != nil {
		log.Fatal(err)
	}

	ircobj.AddCallback("001", func(e *irc.Event) {
		ircobj.Join(conf.Channel)
	})

	ircobj.AddCallback("PRIVMSG", func(e *irc.Event) {
		msgs := strings.Split(string(e.Message()), " ")

		if len(msgs) >= 2 && strings.Contains(msgs[0], "gopher") {
			ircobj.Notice(conf.Channel, prefix+msgs[1]+" not ybsk")
			cmd := &commands.Command{}
			result, err := cmd.Call(msgs[1:])
			if err != nil {
				ircobj.Privmsg(conf.Channel, prefix+fmt.Sprintln(err))
			} else {
				ircobj.Privmsg(conf.Channel, result)
			}
		}
	})

	ircobj.Loop()
}
