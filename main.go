package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/B.E.A.T/command"
	"github.com/B.E.A.T/helper"
	"github.com/joho/godotenv"
	"gopkg.in/tucnak/telebot.v2"
)

var env, token string

func init() {
	if env := godotenv.Load(); env != nil {
		log.Panicln(fmt.Errorf("REASON: %s", env))
	}
}

func getEnv() string {
	env = os.Getenv("RUN_ENV")

	return env
}

func getToken() {
	getEnv()

	if env == "local" {
		token = os.Getenv("TOKEN_LOCAL")
	} else if env == "staging" {
		token = os.Getenv("TOKEN_STAGING")
	}
}

func main() {
	getToken()

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	helper.ErrorMessage(err)

	bot.Handle(telebot.OnText, func(m *telebot.Message) {
		command.Main(bot, m)
	})

	bot.Start()
}
