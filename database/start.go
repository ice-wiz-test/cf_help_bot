package database

import (
	"context"
	"database/sql"
	"log"
	"os"
	"strconv"
)

// TODO - should change database/sql to pgx
var haveOpenConnection bool = false
var openedConnection *sql.DB

func get_connection_string() string {
	filePath := "database/startup/connection_string.txt"
	readFile, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(readFile)
}

func open_connection() (bool, error) {
	if haveOpenConnection {
		return true, nil
	}
	s1 := get_connection_string()
	db, err := sql.Open("postgres", s1)
	if err != nil {
		return false, err
	}
	haveOpenConnection = true
	openedConnection = db
	return true, nil
}

func Does_person_exist_in_database_by_UserID(id int) (error, bool) {
	_, err := open_connection()
	if err != nil {
		return err, false
	}
	query_string := "SELECT * FROM telegram_bot WHERE telegram_bot.userID = "
	query_string += strconv.Itoa(id)
	res, er := openedConnection.QueryContext(context.Background(), query_string)
	if er != nil {
		return nil, false
	}
	return nil, res.Next()
}

func Does_person_exist_in_database_by_handle(handle string) (error, bool) {
	_, err := open_connection()
	if err != nil {
		return err, false
	}
	query_string := "SELECT * FROM telegram_bot WHERE telegram_bot.handle = "
	query_string += handle
	res, er := openedConnection.QueryContext(context.Background(), query_string)
	log.Println(res)
	if er != nil {
		return nil, false
	}
	return nil, res.Next()
}
