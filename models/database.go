package models

import (
	"database/sql"
	"log"
)

func GetMysql() *sql.DB {
	// sqlUser,sqlPass,sqlNameは別ファイルにて管理(gitでは管理外)
	db, err := sql.Open("mysql", sqlUser+":"+sqlPass+"@/"+sqlName)
	if err != nil {
		log.Fatal("エラー：", err)
		// *sql.DB型の何かを返す
	}
	return db
}
