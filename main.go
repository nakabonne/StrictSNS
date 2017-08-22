package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	defer db.Close()
	http.HandleFunc("/", userIndex)
	http.HandleFunc("/post/new", userNew)
	http.HandleFunc("/post/create", userCreate)
	http.HandleFunc("/post/detail/", userDetail)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", Log(http.DefaultServeMux))
	}
}
