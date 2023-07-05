package bot

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var lang string

func welcome(update tgbotapi.Update, bot tgbotapi.BotAPI) {
	switch lang {
	case "Eng":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello! Welcome to the CFHelpBot!")
		send(bot, msg)
		// Send description message
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "")
	case "Rus":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Это CFHelpBot!")
		send(bot, msg)
		// Send description message
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "")
	}
}

func send(bot tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	_, sendErr := bot.Send(msg)
	if sendErr != nil {
		log.Println(sendErr)
	}
}

func Initialize() {
	// Create a new bot instance
	// Get token from the hidden file
	data, err := os.ReadFile("bot/startup/token.txt")
	if err != nil {
		log.Fatal(err)
	}

	token := string(data)
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Println("bot")
		log.Fatal(err)
	}

	// Enable debug mode
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Set up an update configuration
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	// Get updates from the bot
	updateChan, err := bot.GetUpdatesChan(updateConfig)

	if err != nil {
		log.Fatal(err)
	}
	// Process received updates
	for update := range updateChan {
		if update.Message == nil {
			// Ignore any non-Message updates
			continue
		}

		// Check if the message contains "/start"
		if strings.Contains(update.Message.Text, "/start") {
			// Send message to start language select
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please select your language(/lang)")
			send(*bot, msg)
		}

		if strings.Contains(update.Message.Text, "/lang") {
			// Send language select message
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "What language do you want. Please write langEng or langRus")
			send(*bot, msg)
		}

		if strings.Contains(update.Message.Text, "langEng") {
			lang = "Eng"
			// Send a welcome message
			welcome(update, *bot)

		}

		if strings.Contains(update.Message.Text, "langRus") {
			lang = "Rus"
			// Send a welcome message
			welcome(update, *bot)
		}
	}

}
