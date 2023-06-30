package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	user "cf_help_bot/user"
	api "cf_help_bot/api"

 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	u := user.User{}
	u.Handle = "LeftPepeper"
	data := api.getUserRating(u)
	fmt.Println(data)
	// Create a new bot instance
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))

	if err != nil {
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
		if update.Message == nil { // Ignore any non-Message updates
			continue
		}

 		// Check if the message contains "/start"
		if strings.Contains(update.Message.Text, "/start") {
			// Send a welcome message
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello! Welcome to the bot!")
			_, sendErr := bot.Send(msg)
			if sendErr != nil {
				log.Println(sendErr)
			}
		}
	}
	
}