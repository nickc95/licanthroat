package main

import (
	"errors"
)

// TODO: Store all active game channels locally
type gameChannels struct {
	channels []gameChannel
}

type gameChannel struct {
	channelID string
	session   *gameSession
}

func initGameChannels() *gameChannels {
	return &gameChannels{}
}

func (gameChannels *gameChannels) isActive(channelID string) bool {
	for _, channel := range gameChannels.channels {
		if channel.channelID == channelID {
			return true
		}
	}

	return false
}

func (gameChannels *gameChannels) add(channelID string) error {
	// TODO: no duplicate channels
	newGameChannel := gameChannel{
		channelID: channelID,
	}
	gameChannels.channels = append(gameChannels.channels, newGameChannel)
	return nil
}

func (gameChannels *gameChannels) remove(channelID string) error {
	for i, channel := range gameChannels.channels {
		if channel.channelID == channelID {
			gameChannels.channels[i] = gameChannels.channels[len(gameChannels.channels)-1]
			gameChannels.channels = gameChannels.channels[:len(gameChannels.channels)-1]
			return nil
		}
	}

	return errors.New("channel corresponding to given argument was not found")
}
