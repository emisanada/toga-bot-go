package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/bwmarrin/discordgo"
	"github.com/emisanada/toga-bot-go/pkg/exchange"
)

func init() {
	flag.StringVar(&token, "t", "", "Discord Authentication Token")
	flag.Parse()
}

var token string
var buffer = make([][]byte, 0)

func main() {

	if token == "" {
		log.Warn("Please provide an authentication token")
		return
	}

	tg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Error opening Discord connection")
		return
	}

	tg.AddHandler(messageCreate)

	err = tg.Open()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Error opening Discord session")
	}

	log.Info("Starting Toga Bot")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	tg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "TOGA" {
		s.ChannelMessageSend(m.ChannelID, "LUSY")
	}

	if m.Content == "LUSY" {
		s.ChannelMessageSend(m.ChannelID, "TOGA")
	}

	if m.Content == "exchange" {
		s.ChannelMessageSend(m.ChannelID, exchange.GetPrice("coal"))
	}
}
