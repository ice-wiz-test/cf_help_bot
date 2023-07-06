package database

import (
	"context"
	db "github.com/jackc/pgx/v4"
	"log"
)

func get_connection() (error, *db.Conn) {
	conn, err := db.Connect(context.Background(), "postgres://postgres:qCdMMnPsYYt6Ss6AqKeL@localhost:5432/test")
	if err != nil {
		log.Println(err)
	}
	return err, conn
}
