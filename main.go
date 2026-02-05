package main

import (
	"github.com/bopke/korwin"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	session, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}
	session.AddHandler(OnMessageCreate)
	err = session.Open()
	if err != nil {
		panic(err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	err = session.Close()
	if err != nil {
		panic(err)
	}
}

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, mention := range m.Mentions {
		if mention.ID == s.State.User.ID {
			_, _ = s.ChannelMessageSend(m.ChannelID, korwin.GenerateStatement())
			return
		}
	}
}
