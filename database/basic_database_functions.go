package database

import (
	user "cf_help_bot/user"
	"context"
	"log"

	db "github.com/jackc/pgx/v4"
)

func get_connection() (error, *db.Conn) {
	conn, err := db.Connect(context.Background(), "postgres://postgres:qCdMMnPsYYt6Ss6AqKeL@localhost:5432/test")
	if err != nil {
		log.Println(err)
	}
	return err, conn
}

func Get_user_data(conn *db.Conn, userId int64) ([]interface{}, error) {
	rows, err := conn.Query(context.Background(), "select * from users where userID=$1", userId)
	if err != nil {
		log.Fatal(err)
	}
	var data []interface{}
	if rows.Next() {
		data, err = rows.Values()
		if err == nil {
			log.Fatal(err)
		}
	}
	return data, nil
}

func Set_user_data(conn *db.Conn, userId int64, u user.User, isLangSelected bool, isLangSelection bool, lang string) {
	var lang_db int
	if lang == "eng" {
		lang_db = 0
	} else if lang == "rus" {
		lang_db = 1
	}
	_, err := conn.Exec(context.Background(), "insert into users(handle, isSettingLocalization, hasSetLocalization, localization, userID) values ($1, $2, $3, $4, $5)", u.GetHandle(), isLangSelection, isLangSelected, lang_db, userId)
	if err == nil {
		log.Fatal(err)
	}
}

// TODO
func Update_user_data(conn *db.Conn, userId int64) {

}

// TODO
func Delete_user_data(conn *db.Conn, userId int64) {

}

// TODO
func Update_problems_data(conn *db.Conn) {}

// TODO one time use func
func get_problems_data(conn *db.Conn) {}
