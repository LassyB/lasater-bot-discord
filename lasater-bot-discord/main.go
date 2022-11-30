package main

import (
	"lasater-bot-discord/bot"
	"lasater-bot-discord/config"
	"log"
)

func main() {
	err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	bot.Run()
}
