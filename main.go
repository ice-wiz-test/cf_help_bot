package main

import (
	bot "cf_help_bot/bot"
	//database "cf_help_bot/database"
	web "cf_help_bot/web"
	//"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	go bot.Initialize()
	go web.Start()
	select {}
}
