package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	telebot "gopkg.in/telebot.v3"
)

func main() {
	configuration, err := LoadConfiguration()
	if err != nil {
		log.Fatal(err)
		return
	}

	wordleSolver := CreateWordleSovler(configuration.Dictionary)

	bot, err := telebot.NewBot(telebot.Settings{
		Token:  configuration.TelegramBotKey,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle("/echo", func(context telebot.Context) error {
		log.Printf("@%s: /echo %s\n", context.Sender().Username, context.Message().Payload)
		return context.Send(fmt.Sprintf("Echo: %s", context.Message().Payload))
	})

	bot.Handle("/whoami", func(context telebot.Context) error {
		sender := context.Sender()

		log.Printf("@%s: /whoami", sender.Username)

		return context.Send(fmt.Sprintf("ID: %v\nUsername: %s\nName: %s %s",
			sender.ID, sender.Username, sender.FirstName, sender.LastName))
	})

	bot.Handle("/wordle", func(context telebot.Context) error {
		log.Printf("@%s: /wordle %s", context.Sender().Username, context.Message().Payload)

		correct, misplaced, incorrect, err := extractWordleArguments(context.Args())
		if err != nil {
			return context.Send(fmt.Sprintf("Invalid Wordle parameters: %s", err))
		}

		matches, err := wordleSolver.Solve(correct, misplaced, incorrect)

		if err != nil {
			log.Printf("Error solving Wordle: %s", err)
			return context.Send("Error encountered while solving, please try again")
		}

		message := strings.Builder{}
		message.WriteString("Correct Letters: ")
		message.WriteString(correct)
		message.WriteRune('\n')
		message.WriteString("Misplaced Letters: ")
		message.WriteString(misplaced)
		message.WriteRune('\n')
		message.WriteString("Incorrect Letters: ")
		message.WriteString(incorrect)
		message.WriteRune('\n')

		if len(matches) > 0 {
			message.WriteString("Matches:\n")
			message.WriteString(strings.Join(matches, "\n"))
		} else {
			message.WriteString("No matches found")
		}

		return context.Send(message.String())
	})

	bot.Start()
}

func extractWordleArguments(arguments []string) (string, string, string, error) {
	if len(arguments) == 0 || arguments[0] == "" {
		return "", "", "", fmt.Errorf("missing \"correct\" parameter")
	}

	correct := strings.ReplaceAll(strings.ToLower(arguments[0]), ".", "_")
	misplaced := wordleArgumentAt(1, arguments)
	incorrect := wordleArgumentAt(2, arguments)

	return correct, misplaced, incorrect, nil
}

func wordleArgumentAt(position int, arguments []string) string {
	if position >= len(arguments) {
		return ""
	} else if argument := arguments[position]; argument == "\"\"" || argument == "_" || argument == "." {
		return ""
	} else {
		return strings.ToLower(argument)
	}
}
