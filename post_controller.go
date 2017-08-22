package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"./models"
)

var db = models.GetMysql()

func userNew(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("views/posts/new.tmpl"))
	if err := temp.ExecuteTemplate(w, "new.tmpl", nil); err != nil {
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
	posts := models.AllPosts(db)
	for _, v := range posts {
		fmt.Println(v)
	}
	temp := template.Must(template.ParseFiles("views/posts/index.tmpl"))
	err := temp.Execute(w, posts)
	if err != nil {
		log.Fatal("テンプレートエラー", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//posts := models.AllPosts(db)
	//defer db.Close()
	//defer posts.Close()
}

func userDetail(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	fmt.Println("idは", segs[3])
	id, _ := strconv.ParseUint(segs[3], 10, 64)
	//for _, v := range segs {
	//	fmt.Println(v)
	//}
	post := models.PostByID(db, id)
	tmpl := template.Must(template.ParseFiles("views/posts/detail.tmpl"))
	err := tmpl.Execute(w, post)
	if err != nil {
		log.Fatal("テンプレートエラー", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
