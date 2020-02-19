package main

type player struct {
	isAlive          bool
	isMarkedForDeath bool
}

func newPlayer() *player {
	newPlayer := &player{
		isAlive: true,
		isMarkedForDeath: false,
	}

	return newPlayer
}
