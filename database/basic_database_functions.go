package database

import (
	user "cf_help_bot/user"
	"context"
	"log"
	"strconv"

	db "github.com/jackc/pgx/v4"
)

func get_connection() (error, *db.Conn) {
	conn, err := db.Connect(context.Background(), "postgres://postgres:qCdMMnPsYYt6Ss6AqKeL@localhost:5432/test")
	if err != nil {
		log.Println(err)
	}
	return err, conn
}

func Get_user_data(userId int) ([]interface{}, error) {
	err, person_exists := Does_person_exist_in_database(int(userId))
	if err != nil {
		log.Fatal(err)
	}
	if person_exists {
		query_string := "SELECT * FROM telegram_bot WHERE telegram_bot.userID = "
		query_string += strconv.Itoa(userId)
		rows, err := openedConnection.QueryContext(context.Background(), query_string)
		if err != nil {
			log.Fatal(err)
		}
		var data []interface{}
		for rows.Next() {
			if err := rows.Scan(&data); err != nil {
				log.Fatal(err)
			}
		}
		rows.Close()
		openedConnection.Close()
		return data, nil
	} else {
		return nil, nil
	}
}

func Set_user_data(userId int, u user.User, isLangSelected bool, isLangSelection bool, lang string) {
	_, err := open_connection()
	var lang_db int
	if lang == "eng" {
		lang_db = 0
	} else if lang == "rus" {
		lang_db = 1
	}
	result, err := openedConnection.ExecContext(context.Background(), "insert into users(handle, isSettingLocalization, hasSetLocalization, localization, userID) values ("+u.GetHandle()+","+strconv.FormatBool(isLangSelection)+","+strconv.FormatBool(isLangSelected)+","+strconv.Itoa(lang_db)+","+strconv.Itoa(userId)+")")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
	}
	openedConnection.Close()
}

// TODO
func Update_user_data(userId int) {

}

// TODO
func Delete_user_data(userId int) {

}

// TODO
func Update_problems_data() {}

// TODO one time use func
func get_problems_data() {}
