package main

import (
	bot "cf_help_bot/bot"
	user "cf_help_bot/user"
	web "cf_help_bot/web"
	"fmt"
)

func main() {
	u := user.User{}
	u.Initialize("LeftPepeper")
	fmt.Println(u.Get_solved_quantity_by_tags())
	fmt.Println(u.Get_solved_indexes_by_tags())
	go bot.Initialize()
	go web.Start()
	select {}
}
