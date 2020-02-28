package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	toga "github.com/emisanada/toga-bot-go"
)

var Session, _ = toga.New()

func init() {
	// Discord Authentication Token
	Session.Token = os.Getenv("DG_TOKEN")
	if Session.Token == "" {
		flag.StringVar(&Session.Token, "t", "", "Discord Authentication Token")
	}
}

func main() {
	if Session.Token == "" {
		log.Warn("Please provide an authentication token")
		return
	}

	err := Session.Open()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Error opening Discord connection")
		os.Exit(1)
	}

	log.Info("Starting Toga Bot")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Session.Close()
}
