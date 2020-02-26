package main

type werewolf struct {
	*player
}

func newWerewolf(userChannelID string) *werewolf {
	newWerewolf := &werewolf{
		player: newPlayer(userChannelID),
	}

	return newWerewolf
}
