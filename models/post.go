package models

import (
	"database/sql"
	"log"
)

type Post struct {
	Id   int
	User_id string
  Content string
	//CreatedAt time.Time
	//UpdatedAt time.Time
}

func AllPosts(db *sql.DB) *sql.Rows {
	rows, err := db.Query("SELECT * FROM `users`")
	if err != nil {
		log.Fatal("エラー：", err)
		// なんか返す
	}
	return rows
}