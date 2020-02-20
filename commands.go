package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func handleCommands(session *discordgo.Session, guild *discordgo.Guild, channel *discordgo.Channel, messageContent string) {
	if strings.HasPrefix(messageContent, "!commands") {
		helpCommand(session, channel.ID)
		return
	}

	if strings.HasPrefix(messageContent, "!init") {
		initCommand(session, channel.ID, guild.Members)
		return
	}

	if strings.HasPrefix(messageContent, "!reset") {
		resetCommand(session, channel.ID)
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

func initCommand(session *discordgo.Session, channelID string, members []*discordgo.Member) {
	if activeChannels.isActive(channelID) == true {
		_, err := session.ChannelMessageSend(channelID, gameAlreadyInSessionMessage)
		if err != nil {
			fmt.Println(err)
			return
		}

		return
	}

	userIDList := []string{}
	for _, v := range members {
		// don't add the bot itself
		if v.User.ID != session.State.User.ID {
			userIDList = append(userIDList, v.User.ID)
		}
	}

	err := activeChannels.newChannel(channelID, userIDList)
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
		err := activeChannels.removeChannel(channelID)
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
