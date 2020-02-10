package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func handleCommands(session *discordgo.Session, messageContent string, channelID string) {
	if strings.HasPrefix(messageContent, "!commands") {
		helpCommand(session, channelID)
		return
	}

	if strings.HasPrefix(messageContent, "!init") {
		initCommand(session, channelID)
		return
	}

	if strings.HasPrefix(messageContent, "!reset") {
		resetCommand(session, channelID)
		return
	}
}

func helpCommand(session *discordgo.Session, channelID string) {
	_, err := session.ChannelMessageSend(channelID, commandsMessage)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func initCommand(session *discordgo.Session, channelID string) {
	if activeChannels.isActive(channelID) == true {
		_, err := session.ChannelMessageSend(channelID, gameAlreadyInSessionMessage)
		if err != nil {
			fmt.Println(err)
			return
		}

		return
	}

	err := activeChannels.add(channelID)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = session.ChannelMessageSend(channelID, channelInitMessage)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func resetCommand(session *discordgo.Session, channelID string) {
	if activeChannels.isActive(channelID) == true {
		err := activeChannels.remove(channelID)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = session.ChannelMessageSend(channelID, channelResetMessage)
		if err != nil {
			fmt.Println(err)
			return
		}

		return
	}

	_, err := session.ChannelMessageSend(channelID, nonActiveChannelMessage)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
