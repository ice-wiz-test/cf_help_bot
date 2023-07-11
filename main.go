package main

import (
	bot "cf_help_bot/bot"
	database "cf_help_bot/database"
	web "cf_help_bot/web"
	"database/sql"

	_ "github.com/lib/pq"
)

func main() {
	connstr := database.Get_connection_string()
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	go bot.Initialize()
	go web.Start()
	select {}
}
