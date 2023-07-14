package bot

import (
	db "cf_help_bot/database"
	hf "cf_help_bot/help_func"
	user "cf_help_bot/user"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var lang string

func welcome(update tgbotapi.Update, bot tgbotapi.BotAPI) {
	log.Println(lang)
	switch lang {
	// Welcome messages and description by the selected language
	case "eng":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello! Welcome to the CFHelpBot!")
		send(bot, msg)
		// Send description message
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "//TODO")
		send(bot, msg)
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Write down your handle in codeforces.com")
		send(bot, msg)
	case "rus":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Это CFHelpBot!")
		send(bot, msg)
		// Send description message
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "//TODO")
		send(bot, msg)
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Напишите свой хэндл с codeforces.com")
		send(bot, msg)
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
	lang = "eng"
	isLangSelected := false
	isLangSelection := false
	isUserInitialized := false
	isUserInitialization := false
	mistakes_counter := -1
	for update := range updateChan {
		if update.Message == nil {
			// Ignore any non-Message updates
			continue
		}

		// Check if the message contains "/start"
		if strings.Contains(update.Message.Text, "/start") {
			log.Println("Start")
			// Send message to start language select
			err, person_exits := db.Does_person_exist_in_database(update.Message.From.ID)
			log.Println("Person exists:")
			log.Println(person_exits)
			data, err := db.Get_user_data(update.Message.From.ID)
			log.Println("Database connected")
			log.Println(data)
			if err != nil {
				log.Fatal(err)
			}
			//log.Println(data)
			if isLangSelected == false {
				isLangSelection = true
			} else {
				log.Println("Lang selected")
				if isUserInitialized == false {
					welcome(update, *bot)
					isUserInitialization = true
				}
			}
		}

		// Manual language selection
		if strings.Contains(update.Message.Text, "/lang") {
			// Send language select message
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "What language do you want. Please write Eng or Rus")
			send(*bot, msg)
			// Select language
			isLangSelection = true
			if isUserInitialized == false {
				// Send a welcome message
				welcome(update, *bot)
			}
		}

		if isLangSelection == true {
			if update.Message.Text == "Eng" || update.Message.Text == "eng" ||
				update.Message.Text == "english" || update.Message.Text == "English" {
				lang = "eng"
				log.Println("Selected lang is Eng")
				isLangSelected = true
				isLangSelection = false
				if isUserInitialized == false {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Log in to your account or register.\n/login")
					send(*bot, msg)
				}
				continue
			} else if update.Message.Text == "Rus" || update.Message.Text == "rus" ||
				update.Message.Text == "русский" || update.Message.Text == "Русский" ||
				update.Message.Text == "рус" || update.Message.Text == "Рус" {
				lang = "rus"
				isLangSelected = true
				isLangSelection = false
				if isUserInitialized == false {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Войдите в аккаунт или зарегистрируйтесь. /login")
					send(*bot, msg)
				}
				continue
			} else {
				if mistakes_counter < 0 {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please select your language. Write Eng or Rus")
					send(*bot, msg)
					mistakes_counter += 1
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You need to write down language you need. English(Eng) or Russian(Rus)")
					send(*bot, msg)
				}
			}
		}

		if isUserInitialized == false {
			if strings.Contains(update.Message.Text, "/login") {
				welcome(update, *bot)
				log.Println(update.Message.Text)
				isUserInitialization = true
			}
		}
		if isUserInitialization {
			if update.Message != nil && update.Message.Text != "/login" {
				log.Println("Message text:")
				log.Println(update.Message.Text)
				u := user.User{}
				// Initialize user by the handle got from user
				u.Initialize(update.Message.From.ID, update.Message.Text, lang, isLangSelected, isLangSelection)
				isUserInitialized = true
				log.Println(u.GetHandle())
				log.Println(u.Get_solved_quantity_by_tags())
				log.Println(u.Get_solved_indexes_by_tags())
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, hf.ConvertMaptToString(u.Get_solved_quantity_by_tags()))
				send(*bot, msg)
				isUserInitialization = false
				db.Set_user_data(update.Message.From.ID, u, u.IsLangSelected(), u.IsLangSelection(), u.GetLang())
			}
		}
	}
}
