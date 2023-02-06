package main

import (
	"fmt"
	"os"
)

const defaultDictionary string = "/usr/share/dict/american-english"

type Configuration struct {
	TelegramBotKey string
	Dictionary     string
}

func LoadConfiguration() (*Configuration, error) {
	botKey := os.Getenv("TELEGRAM_BOT_KEY")
	if botKey == "" {
		return nil, fmt.Errorf("bot key (TELEGRAM_BOT_KEY) is required but not set")
	}

	dictionary := os.Getenv("DICTIONARY")
	if dictionary == "" {
		dictionary = defaultDictionary
	}

	return &Configuration{
		TelegramBotKey: botKey,
		Dictionary:     dictionary,
	}, nil
}
