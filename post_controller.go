package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"./models"
)

func userIndex(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("views/posts/index.html"))
	err := temp.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatal("[golang server] internal server error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println("インデックス")
	db := models.GetMysql()
	posts := models.AllPosts(db)
	defer db.Close()
	defer posts.Close()
}
