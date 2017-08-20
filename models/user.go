package models

import (
	"database/sql"
	"log"
)

type User struct {
	Id   int
	Name string
}

func AllUsers(db *sql.DB) *sql.Rows {
	rows, err := db.Query("SELECT * FROM `users`")
	if err != nil {
		log.Fatal("エラー：", err)
		// なんか返す
	}
	return rows
}
