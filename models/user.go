package models

import (
	"database/sql"
	"log"
)

type User struct {
	ID   int
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

// UserByName id指定で全件取得
func UserByName(db *sql.DB, name string) *User {
	query := "SELECT * FROM users WHERE name = '" + name + "'"
	log.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("クエリーエラー：", err)
	}

	var user *User

	for rows.Next() {
		var (
			id   int
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal("スキャンエラー: ", err)
		}
		user = &User{ID: id, Name: name}
	}
	rows.Close()
	return user
}
