package main

import (
	"log"
	"net/http"

	"./models"
	_ "github.com/go-sql-driver/mysql"
)

var db = models.GetMysql()

func main() {
	defer db.Close()
	http.HandleFunc("/login", userLogin)
	http.HandleFunc("/signin", userSignIn)
	http.HandleFunc("/", postIndex)
	http.HandleFunc("/post/new", postNew)
	http.HandleFunc("/post/create", postCreate)
	http.HandleFunc("/post/detail/", postDetail)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", Log(http.DefaultServeMux))
	}
}
