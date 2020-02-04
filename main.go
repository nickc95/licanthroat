package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func init() {
	configInit()
}

func main() {
	_, err := discordgo.New("Bot " + envConfig.BotToken)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
