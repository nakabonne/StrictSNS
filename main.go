package main

import (
	"fmt"

	"./models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := models.GetMysql()
	posts := models.AllPosts(db)
	defer db.Close()
	defer posts.Close()
	fmt.Println(posts)
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
