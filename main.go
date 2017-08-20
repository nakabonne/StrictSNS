package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", userIndex)
	http.ListenAndServe(":8080", nil)
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
