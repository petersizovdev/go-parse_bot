package main

import (
	"log"

	"github.com/petersizovdev/go-parse_bot.git/internal/clients/tg"
	"github.com/petersizovdev/go-parse_bot.git/internal/config"
	"github.com/petersizovdev/go-parse_bot.git/internal/model/messages"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}

	tgClient, err := tg.New(config)
	if err != nil {
		log.Fatal("tg client init failed:", err)
	}

	msgModel := messages.New(tgClient)

	tgClient.ListenUpdates(msgModel)
}
