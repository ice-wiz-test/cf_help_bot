package main

import (
	bot "cf_help_bot/bot"
	web "cf_help_bot/web"
	_ "github.com/lib/pq"
)

func main() {
	go bot.Initialize()
	go web.Start()
	select {}
}
