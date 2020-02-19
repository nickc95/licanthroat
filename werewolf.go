package main

type werewolf struct {
	*player
}

func newWerewolf() *werewolf {
	newWerewolf := &werewolf{
		player: newPlayer(),
	}

	return newWerewolf
}
