package main

type gameSession struct {
	players     map[string]player
	isNightTime bool
}

func newGameSession() *gameSession {
	gameSession := gameSession{}

	return &gameSession
}
