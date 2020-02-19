package main

type villager struct {
	*player
}

func newVillager() *villager {
	newVillager := &villager{
		player: newPlayer(),
	}

	return newVillager
}
