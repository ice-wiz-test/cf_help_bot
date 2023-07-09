package main

import (
	bot "cf_help_bot/bot"
	web "cf_help_bot/web"
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func get_connection_string() string {
	filePath := "connection_string.txt"
	readFile, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(readFile)
}

func main() {
	connstr := get_connection_string()
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	go bot.Initialize()
	go web.Start()
	select {}
}
