package main

type gameSession struct {
	villagers   map[string]villager
	werewolves  map[string]werewolf
	isNightTime bool
}

func newGameSession() *gameSession {
	return &gameSession{
		isNightTime: true,
		villagers:   make(map[string]villager),
		werewolves:  make(map[string]werewolf),
	}
}

func (gameSession *gameSession) addVillager(playerID string) error {
	gameSession.villagers[playerID] = *newVillager()

	return nil
}

func (gameSession *gameSession) addWerewolf(playerID string) error {
	gameSession.werewolves[playerID] = *newWerewolf()

	return nil
}
