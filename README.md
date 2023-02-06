# cbojar.telegrambot

A simple Telegram bot that does very simple things, like echo back a message, describe the user, and solve Wordle
puzzles. It is built with the telebot SDK for Telegram bots.

To build, run:

```sh
go build ./... && go build -o telegrambot
```

To run, do:
```sh
export TELEGRAM_BOT_KEY='{your Telegram bot token}'
./telegrambot
```

By default, the Wordle solver of the bot will use the dictionary located at `/usr/share/dict/american-english`. To use a
different dictionary, set the `DICTIONARY` environment variable to the path of an alternative dictionary file.

To build in Docker, run:

```sh
docker build -t telegrambot:latest .
```

To run in Docker, do:

```sh
docker run -d --restart=unless-stopped --name telegrambot -e TELEGRAM_BOT_KEY='{your Telegram bot token}' telegrambot:latest
```
