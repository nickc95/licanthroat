package main

type player struct {
	userChannelID    string
	isAlive          bool
	isMarkedForDeath bool
}

func newPlayer(userChannelID string) *player {
	newPlayer := &player{
		userChannelID:    userChannelID,
		isAlive:          true,
		isMarkedForDeath: false,
	}

	return newPlayer
}

func (player *player) getUserChannelID() string {
	return player.userChannelID
}
