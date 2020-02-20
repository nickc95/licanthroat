package main

import (
	"errors"
	"math/rand"
	"time"
)

type gameSession struct {
	villagers   map[string]villager
	werewolves  map[string]werewolf
	isNightTime bool
}

var numberOfWerewolves = 1

func newGameSession(userIDs []string) (*gameSession, error) {
	if len(userIDs) < 3 {
		return nil, errors.New("not enough users to start new session, need at least 3")
	}

	newGameSession := &gameSession{
		isNightTime: true,
		villagers:   make(map[string]villager),
		werewolves:  make(map[string]werewolf),
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numberOfWerewolves; i++ {
		randomIndex := rand.Intn(len(userIDs))
		newGameSession.addWerewolf(userIDs[randomIndex])

		userIDs[randomIndex] = userIDs[len(userIDs)-1]
		userIDs = userIDs[:len(userIDs)-1]
	}
	for _, id := range userIDs {
		newGameSession.addVillager(id)
	}

	return newGameSession, nil
}

func (gameSession *gameSession) addVillager(playerID string) error {
	gameSession.villagers[playerID] = *newVillager()

	return nil
}

func (gameSession *gameSession) addWerewolf(playerID string) error {
	gameSession.werewolves[playerID] = *newWerewolf()

	return nil
}
