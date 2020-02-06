package main

var (
	nonActiveChannelMessage = `
		There is no active game session found for this channel. Please use the !init command to start a new game session.
	`
	commandsMessage = `
		List of supported commands:
		!commands - Get list of supported commands
		!sessionstatus - Get current session status
		!init - Initialize a new game session in this channel
		!reset - Reset game session in this channel
	`
)
