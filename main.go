package main

import (
	"fmt"

	"./models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := models.GetMysql()
	users := models.AllUser(db)
	defer db.Close()
	defer users.Close()
	fmt.Println(users)
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
}
