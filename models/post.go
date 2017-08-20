package models

import (
	"database/sql"
	"log"
)

// Post 投稿
type Post struct {
	ID      int
	UserID  int
	Content string
	//CreatedAt time.Time
	//UpdatedAt time.Time
}

// AllPosts 全件取得
func AllPosts(db *sql.DB) *sql.Rows {
	rows, err := db.Query("SELECT * FROM `posts`")
	if err != nil {
		log.Fatal("エラー：", err)
		// なんか返す
	}
	return rows
}

// Insert インサートする
func (p *Post) Insert(db *sql.DB) {
	query := "INSERT INTO posts (user_id, content) values(?, ?)"
	if _, err := db.Exec(query, p.UserID, p.Content); err != nil {
		log.Fatal("インサートエラー：", err)
	}
}
