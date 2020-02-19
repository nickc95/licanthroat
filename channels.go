package main

import (
	"errors"
)

// TODO: Store all active game channels locally
type gameChannels struct {
	channels map[string]gameChannel
}

type gameChannel struct {
	channelID string
	session   *gameSession
}

func initGameChannels() *gameChannels {
	return &gameChannels{
		channels: make(map[string]gameChannel),
	}
}

func (gameChannels *gameChannels) isActive(channelID string) bool {
	_, ok := gameChannels.channels[channelID]
	if ok {
		return true
	}

	return false
}

func (gameChannels *gameChannels) add(channelID string) error {
	newGameChannel := gameChannel{
		channelID: channelID,
	}
	gameChannels.channels[channelID] = newGameChannel
	return nil
}

func (gameChannels *gameChannels) remove(channelID string) error {
	_, ok := gameChannels.channels[channelID]
	if ok {
		delete(gameChannels.channels, channelID)
		return nil
	}

	return errors.New("channel corresponding to given argument was not found")
}
