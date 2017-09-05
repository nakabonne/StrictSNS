package models

import (
	"database/sql"
	"log"
	"strconv"
	"time"
)

// Post 投稿
type Post struct {
	ID        int
	UserID    int
	Content   string
	CreatedAt []uint8
	UpdatedAt []uint8
}

// Insert インサートする
func (p *Post) Insert(db *sql.DB) {
	query := "INSERT INTO posts (id, user_id, content, created_at, updated_at) values(?, ?, ?, ?, ?)"
	if _, err := db.Exec(query, p.ID, p.UserID, p.Content, time.Now(), time.Now()); err != nil {
		log.Fatal("インサートエラー：", err)
	}
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
		posts = append(posts, &Post{ID: id, UserID: userID, Content: content, CreatedAt: createdAt, UpdatedAt: updatedAt})
	}
	rows.Close()
	return posts
}

// PostByID id指定で全件取得
func PostByID(db *sql.DB, id uint64) *Post {
	idS := strconv.FormatUint(id, 10)
	query := "SELECT * FROM `posts` WHERE `id` = " + idS
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("クエリーエラー：", err)
	}

	var post *Post

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
		post = &Post{ID: id, UserID: userID, Content: content, CreatedAt: createdAt, UpdatedAt: updatedAt}
	}
	rows.Close()
	return post
}
