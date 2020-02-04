package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type config struct {
	BotToken string `env:"Token" envDefault:"NjczNzI3MTYxNzUyMDkyNjcz.XjmywQ.BD7j7r6gw4Jfomc4lmQgOfHHJlY"`
}

var envConfig *config

func configInit() {
	envConfig = &config{}
	if err := env.Parse(envConfig); err != nil {
		fmt.Println(err.Error())
	}
}
