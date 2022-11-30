package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"lasater-bot-discord/config"

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
	if !strings.HasPrefix(message.Content, "!lasa") {
		return
	}
	// Respond to messages
	splitCommand := strings.Split(message.Content, " ")
	fmt.Println(splitCommand[1])
	switch splitCommand[1] {
	case "ping":
		discord.ChannelMessageSend(message.ChannelID, "I'm alive")
	case "dnd":
		dndSubcategory := splitCommand[2]
		switch dndSubcategory {
		case "spell":
			discord.ChannelMessageSend(message.ChannelID, strings.Join(splitCommand[3:], " "))
		}
	}
}
