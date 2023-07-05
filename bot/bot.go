package bot

import (
	user "cf_help_bot/user"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var lang string

func welcome(update tgbotapi.Update, bot tgbotapi.BotAPI) {
	switch lang {
	// Welcome messages and description by the selected language
	case "Eng":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello! Welcome to the CFHelpBot!")
		send(bot, msg)
		// Send description message
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "//TODO")
		send(bot, msg)
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Write down your handle in codeforces.com")
		send(bot, msg)
	case "Rus":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Это CFHelpBot!")
		send(bot, msg)
		// Send description message
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "//TODO")
		send(bot, msg)
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Напишите свой хэндл с codeforces.com")
		send(bot, msg)
	}
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
		} else {
			log.Println(update.Message.Text)
			u := user.User{}
			// Initialize user by the handle got from user
			u.Initialize(update.Message.Text)
			log.Println(u.GetHandle())
			log.Println(u.Get_solved_quantity_by_tags())
			log.Println(u.Get_solved_indexes_by_tags())
		}
	}
}

func send(bot tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {
	_, sendErr := bot.Send(msg)
	if sendErr != nil {
		log.Println(sendErr)
	}
}

func selectLang(bot tgbotapi.BotAPI) {
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
		// Select language
		if update.Message.Text == "Eng" || update.Message.Text == "eng" ||
			update.Message.Text == "english" || update.Message.Text == "English" {
			lang = "Eng"
			return
		}
		if update.Message.Text == "Rus" || update.Message.Text == "rus" ||
			update.Message.Text == "русский" || update.Message.Text == "Русский" ||
			update.Message.Text == "рус" || update.Message.Text == "Рус" {
			lang = "Rus"
			return
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You need to write down language you need. English(Eng) or Russian(Rus)")
			send(bot, msg)
		}
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
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "What language do you want. Please write Eng or Rus")
			send(*bot, msg)
			// Select language
			selectLang(*bot)
			// Send a welcome message
			welcome(update, *bot)
		}
	}

}
