package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func init() {
	configInit()
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
	fmt.Println("Ready")

	// TODO: leave all guilds bot is a part of (state reset)
}

// sent when a message is created, m is a message struct
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore messages created by itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// TODO: Refactor out command handlers
	if strings.HasPrefix(m.Content, "!help") {
		// TODO: check that channel message came from is an 'active game channel'
		channel, err := s.State.Channel(m.ChannelID)
		if err != nil {
			fmt.Println("couldn't find corresponding channel")
			return
		}

		// guild, err = s.State.Guild(channel.GuildID)
		// if err != nil {
		// 	fmt.Println("couldn't find corresponding guild")
		// 	return
		// }

		_, _ = s.ChannelMessageSend(channel.ID, "i'm sending a message corresponding to the message !help")
		return
	}
}

// sent when a new guild is joined
func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}

	// TODO: Create a new channel in guild specific for game

	// for _, channel := range event.Guild.Channels {
	// 	if channel.ID == event.Guild.ID {
	// 		_, _ = s.ChannelMessageSend(channel.ID, "licanthroat is ready! Type nothing because nothing is supported yet.")
	// 		return
	// 	}
	// }
}
