package main

import (
	bot "cf_help_bot/bot"
	db "cf_help_bot/database"
	web "cf_help_bot/web"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	go bot.Initialize()
	go web.Start()
	go log.Println(db.Does_person_exist_in_database_by_handle("LeftPepeper"))
	select {}
}
