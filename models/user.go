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

/* 使い方
for users.Next() {
	var id int
	var name string
	if err := users.Scan(&id, &name); err != nil {
		log.Fatal("エラー: ", err)
	}
	fmt.Println(id, name)
}
*/
