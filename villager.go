package main

type villager struct {
	*player
}

func newVillager(userChannelID string) *villager {
	newVillager := &villager{
		player: newPlayer(userChannelID),
	}

	return newVillager
}
