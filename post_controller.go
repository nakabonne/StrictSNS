package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"./models"
)

func userNew(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("views/posts/new.html"))
	if err := temp.ExecuteTemplate(w, "new.html", nil); err != nil {
		log.Fatal("[golang server] internal server error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func userCreate(w http.ResponseWriter, r *http.Request) {
	// リクエストをパース
	if err := r.ParseForm(); err != nil {
		log.Fatal("エラー：", err)
	}
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	http.Redirect(w, r, "/", 301)
}

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
