package models

import (
	"database/sql"
	"log"
)

// Post 投稿
type Post struct {
	ID        int
	UserID    int
	Content   string
	CreatedAt []uint8
	UpdatedAt []uint8
}

// AllPosts 全件取得
func AllPosts(db *sql.DB) []*Post {
	rows, err := db.Query("SELECT * FROM `posts`")
	if err != nil {
		log.Fatal("クエリーエラー：", err)
		// なんか返す
	}

	posts := []*Post{}

	for rows.Next() {
		var (
			id        int
			userID    int
			content   string
			createdAt []uint8
			updatedAt []uint8
		)
		if err := rows.Scan(&id, &userID, &content, &createdAt, &updatedAt); err != nil {
			log.Fatal("スキャンエラー: ", err)
		}
		posts = append(posts, &Post{ID: id, UserID: userID, Content: content})
	}
	return posts
}

// Insert インサートする
func (p *Post) Insert(db *sql.DB) {
	query := "INSERT INTO posts (user_id, content) values(?, ?)"
	if _, err := db.Exec(query, p.UserID, p.Content); err != nil {
		log.Fatal("インサートエラー：", err)
	}
}
