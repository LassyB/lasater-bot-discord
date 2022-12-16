package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"lasater-bot-discord/config"
	"lasater-bot-discord/dnd"

	"github.com/bwmarrin/discordgo"
)

// Store Bot API Tokens:
var (
	BotToken string
)

func Run() {
	// Create new Discord Session
	discord, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
	}

	// Add event handler for general messages
	discord.AddHandler(newMessage)

	// Open session
	discord.Open()
	defer discord.Close()

	// Run until code is terminated
	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore bot message
	if message.Author.ID == discord.State.User.ID {
		return
	}
	// Respond to messages
	splitCommand := strings.Split(message.Content, " ")
	messageContents := ""
	switch splitCommand[0] {
	case "!ping":
		messageContents = "pong"
	case "!dnd":
		messageContents = dnd.HandleMessage(splitCommand[1], strings.Join(splitCommand[2:], " "))
	}
	if len(messageContents) > 0 {
		discord.ChannelMessageSend(message.ChannelID, "```"+messageContents+"```")
	}
}
