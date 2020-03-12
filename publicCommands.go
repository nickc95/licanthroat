package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func handlePublicCommands(session *discordgo.Session, guild *discordgo.Guild, channel *discordgo.Channel, messageContent string) {
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

	if strings.HasPrefix(messageContent, "!status") {
		statusCommand(session, channel.ID)
		return
	}

	if strings.HasPrefix(messageContent, "!tick") {
		tickCommand(session, channel.ID)
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

	// userInfoMap maps userIDs (identify user) to userChannelIDs (identify their DM channel)
	userInfoMap := make(map[string]string)
	for _, v := range members {
		// don't add the bot itself
		if v.User.ID != session.State.User.ID {
			userChannel, err := session.UserChannelCreate(v.User.ID)
			if err != nil {
				fmt.Println(err)
				return
			}

			userInfoMap[v.User.ID] = userChannel.ID
		}
	}

	err := activeChannels.newChannel(channelID, userInfoMap)
	if err != nil {
		fmt.Println(err)
		return
	}

	message := privateInitMessage + channelID
	for _, userChannelID := range userInfoMap {
		_, err = session.ChannelMessageSend(userChannelID, message)
		if err != nil {
			fmt.Println(err)
			return
		}
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

func statusCommand(session *discordgo.Session, channelID string) {
	if activeChannels.isActive(channelID) == false {
		_, err := session.ChannelMessageSend(channelID, nonActiveChannelMessage)
		if err != nil {
			fmt.Println(err)
			return
		}

		return
	}

	gameSession, err := activeChannels.getGameSession(channelID)
	if err != nil {
		fmt.Println(err)
		return
	}

	isGameOver, winner := gameSession.isGameOver()
	if isGameOver {
		message := gameOverMessage + winner + "."
		_, err = session.ChannelMessageSend(channelID, message)
		if err != nil {
			fmt.Println(err)
			return
		}

		return
	}

	_, err = session.ChannelMessageSend(channelID, gameNotOverMessage)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func tickCommand(session *discordgo.Session, channelID string) {
	if activeChannels.isActive(channelID) == false {
		_, err := session.ChannelMessageSend(channelID, nonActiveChannelMessage)
		if err != nil {
			fmt.Println(err)
			return
		}

		return
	}

	gameSession, err := activeChannels.getGameSession(channelID)
	if err != nil {
		fmt.Println(err)
		return
	}

	isNowNightTime := gameSession.toggleAndCheckIsNightTime()
	if isNowNightTime {
		// send notification, do nighttime things
		return
	}

	// send notification, do daytime things
}
