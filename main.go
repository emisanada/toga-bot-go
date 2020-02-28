package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"github.com/bwmarrin/discordgo"
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
