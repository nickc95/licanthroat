package main

// roles: werewolf, guard, seer, villager

type player struct {
	role       string
	isAlive    bool
	isGuarded  bool
	isExamined bool
}

func initPlayer(role string) *player {
	newPlayer := &player{
		role:       role,
		isAlive:    true,
		isGuarded:  false,
		isExamined: false,
	}

	return newPlayer
}

func (player *player) isWerewolf() bool {
	if player.role == "werewolf" {
		return true
	}

	return false
}

func (player *player) isInnocent() bool {
	if player.role != "werewolf" {
		return true
	}

	return false
}

func (player *player) kill() bool {
	if !player.isAlive {
		return false
	}

	player.isAlive = false
	return true
}

func (player *player) guard() bool {
	if !player.isAlive {
		return false
	}

	player.isGuarded = true
	return true
}

func (player *player) resetModifiers() {
	player.isGuarded = false
	player.isExamined = false
}
