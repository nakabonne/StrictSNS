package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"./models"
)

func userLogin(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("views/users/login.tmpl"))
	if err := temp.Execute(w, nil); err != nil {
		log.Fatal("テンプレートエラー", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func userSignIn(w http.ResponseWriter, r *http.Request) {
	// リクエストをパース
	if err := r.ParseForm(); err != nil {
		log.Fatal("エラー：", err)
	}
	name := strings.Join(r.Form["name"], "")
	user := models.UserByName(db, name)
	if user == nil {
		http.Redirect(w, r, "/login", 301)
	} else {
		http.Redirect(w, r, "/", 301)
	}
}
