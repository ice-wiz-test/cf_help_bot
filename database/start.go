package database

import (
	"context"
	"database/sql"
	"os"
)

// TODO - should change database/sql to pgx
var haveOpenConnection bool = false
var openedConnection *sql.DB

func Get_connection_string() string {
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
	s1 := Get_connection_string()
	db, err := sql.Open("postgres", s1)
	if err != nil {
		return false, err
	}
	haveOpenConnection = true
	openedConnection = db
	return true, nil
}

func Does_person_exist_in_database(id int64) (error, bool) {
	_, err := open_connection()
	if err != nil {
		return err, false
	}
	query_string := "SELECT TOP 1 telegram_bot.userID FROM telegram_bot WHERE telegram_bot.userID = "
	query_string += string(id)
	res, er := openedConnection.QueryContext(context.Background(), query_string)
	if er != nil {
		return er, false
	}
	return nil, res.Next()
}
