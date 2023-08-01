package database

import (
	user "cf_help_bot/user"
	"context"
	"log"
	"strconv"

	db "github.com/jackc/pgx/v4"
)

type SendedArgs struct {
	userID          int
	handle          string
	isLangSelected  bool
	isLangSelection bool
	lang            string
}

func get_connection() (error, *db.Conn) {
	conn, err := db.Connect(context.Background(), "postgres://postgres:qCdMMnPsYYt6Ss6AqKeL@localhost:5432/test")
	if err != nil {
		log.Println(err)
	}
	return err, conn
}

func Get_user_data_by_UserID(userId int) ([]interface{}, error) {
	err, person_exists := Does_person_exist_in_database_by_UserID(int(userId))
	if err != nil {
		log.Println(err)
	}
	if person_exists {
		query_string := "SELECT * FROM telegram_bot WHERE telegram_bot.userID = "
		query_string += strconv.Itoa(userId)
		rows, err := openedConnection.QueryContext(context.Background(), query_string)
		log.Println("Get query done")
		if err != nil {
			log.Println(err)
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

func Get_user_data_by_handle(handle string) ([]interface{}, error) {
	err, person_exists := Does_person_exist_in_database_by_handle(handle)
	if err != nil {
		log.Println(err)
	}
	if person_exists {
		query_string := "SELECT * FROM telegram_bot WHERE telegram_bot.handle = "
		query_string += handle
		rows, err := openedConnection.QueryContext(context.Background(), query_string)
		log.Println("Get query done")
		if err != nil {
			log.Println(err)
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
	isLangSelectedDB := 0
	isLangSelectionDB := 0
	if isLangSelected == true {
		isLangSelectedDB = 1
	} else {
		isLangSelectedDB = 0
	}
	if isLangSelection == true {
		isLangSelectionDB = 1
	} else {
		isLangSelectionDB = 0
	}
	query := "INSERT INTO telegram_bot (handle, isSettingLocalization, hasSetLocalization, localization, userID) VALUES ($1, $2, $3, $4, $5)"
	result, err := openedConnection.ExecContext(context.Background(), query, u.GetHandle(), isLangSelectionDB, isLangSelectedDB, lang_db, userId)
	if err != nil {
		log.Println(err)
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
func Update_user_data(args SendedArgs) {
	_, err := open_connection()
	if err != nil {
		log.Fatal(err)
	}
}

// Function delete all user data from database
func Delete_user_data(userId int) {
	_, err := open_connection()
	if err != nil {
		log.Fatal(err)
	}
	query := "DELETE FROM telegram_bot WHERE telegram_bot.userID = " + strconv.Itoa(userId)
	result, err := openedConnection.ExecContext(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatalf("expected single row affected, got %d rows affected", rows)
	}
}

// TODO (this func should work on server)
func Update_problems_data() {}

// TODO one time use func
func get_problems_data() {}
