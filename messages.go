package main

var (
	// public messages
	nonActiveChannelMessage = `
		There is no active game session found for this channel. Please use the !init command to start a new game session.
	`
	commandsMessage = `
		List of supported commands:
		!commands - Get list of supported commands
		!status - Get current session status
		!init - Initialize a new game session in this channel
		!reset - Reset game session in this channel
	`

	gameAlreadyInSessionMessage = `
		There is already an active game session for this channel. Please use the !reset command to clear this game session.
	`

	channelInitMessage = `
		Game session initialized.
	`

	channelResetMessage = `
		Game session for this channel was reset. Please use the !init command to start a new game session.
	`

	gameOverMessage = `
		Game is over. Winner is: 
	`

	gameNotOverMessage = `
		Game is not over, there is still either a villager or werewolf still alive!
	`

	// DM messages
	privateInitMessage = `
		Game session has started! This DM channel is linked to session identified by ID: 
	`
)
