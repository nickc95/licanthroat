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

func newGameSession(userInfoMap map[string]string) (*gameSession, error) {
	if len(userInfoMap) < 3 {
		return nil, errors.New("not enough users to start new session, need at least 3")
	}

	userIDPool := []string{}
	for userID := range userInfoMap {
		userIDPool = append(userIDPool, userID)
	}

	newGameSession := &gameSession{
		isNightTime: true,
		villagers:   make(map[string]villager),
		werewolves:  make(map[string]werewolf),
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numberOfWerewolves; i++ {
		randomIndex := rand.Intn(len(userIDPool))
		pickedUserID := userIDPool[randomIndex]
		pickedUserChannelID := userInfoMap[pickedUserID]
		newGameSession.addWerewolf(pickedUserID, pickedUserChannelID)

		userIDPool[randomIndex] = userIDPool[len(userIDPool)-1]
		userIDPool = userIDPool[:len(userIDPool)-1]
	}
	for _, id := range userIDPool {
		newGameSession.addVillager(id, userInfoMap[id])
	}

	return newGameSession, nil
}

func (gameSession *gameSession) addVillager(playerID string, userChannelID string) error {
	gameSession.villagers[playerID] = *newVillager(userChannelID)

	return nil
}

func (gameSession *gameSession) getLiveVillagers() map[string]villager {
	liveVillagers := make(map[string]villager)

	for playerID, villager := range gameSession.villagers {
		if villager.player.isAlive {
			liveVillagers[playerID] = villager
		}
	}

	return liveVillagers
}

func (gameSession *gameSession) addWerewolf(playerID string, userChannelID string) error {
	gameSession.werewolves[playerID] = *newWerewolf(userChannelID)

	return nil
}

func (gameSession *gameSession) getLiveWerewolves() map[string]werewolf {
	liveWerewolves := make(map[string]werewolf)

	for playerID, werewolf := range gameSession.werewolves {
		if werewolf.player.isAlive {
			liveWerewolves[playerID] = werewolf
		}
	}

	return liveWerewolves
}

func (gameSession *gameSession) isGameOver() (isGameOver bool, winner string) {
	liveVillagers := gameSession.getLiveVillagers()
	if len(liveVillagers) == 0 {
		return true, "villagers"
	}

	liveWerewolves := gameSession.getLiveWerewolves()
	if len(liveWerewolves) == 0 {
		return true, "werewolves"
	}

	return false, ""
}
