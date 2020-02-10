package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

// global gameChannels storage
var activeChannels *gameChannels

func init() {
	configInit()
	activeChannels = initGameChannels()
}

func main() {
	discord, err := discordgo.New("Bot " + envConfig.BotToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	discord.AddHandler(ready)
	discord.AddHandler(messageCreate)
	discord.AddHandler(guildCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

// sent when client completes initial handshake with gateway for a new session.
func ready(s *discordgo.Session, event *discordgo.Ready) {
}

// sent when a message is created, m is a message struct
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore messages created by itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("couldn't find corresponding channel")
		return
	}

	handleCommands(s, m.Content, channel.ID)
}

// sent when a new guild is joined
func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}

	// for _, channel := range event.Guild.Channels {
	// 	_, err := s.ChannelMessageSend(channel.ID, "licanthroat is ready! Type !commands to see a list of available commands.")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }
}
