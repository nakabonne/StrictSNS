package models

import (
	"database/sql"
	"log"
)

func GetMysql() *sql.DB {
	db, err := sql.Open("mysql", "root:Tsuba2896@/go_samples")
	if err != nil {
		log.Fatal("エラー：", err)
		// *sql.DB型の何かを返す
	}
	return db
}
