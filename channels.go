package main

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

func (gameChannels *gameChannels) isChannelActive(channelID string) bool {
	for _, channel := range gameChannels.channels {
		if channel.channelID == channelID {
			return true
		}
	}

	return false
}
