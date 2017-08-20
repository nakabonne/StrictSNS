package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"./models"
)

var db = models.GetMysql()

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
	fmt.Println(strings.Join(r.Form["content"], ""))

	content := strings.Join(r.Form["content"], "")
	post := models.Post{
		UserID:  1,
		Content: content,
	}
	post.Insert(db)
	/*
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}
	*/
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
	//posts := models.AllPosts(db)
	//defer db.Close()
	//defer posts.Close()
}
