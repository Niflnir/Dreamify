package database

import (
	"database/sql"
	"fmt"
)

func CreatePost(db *sql.DB, title string, body string) {
	_, err := db.Exec("INSERT INTO posts(title,body) VALUES($1,$2)", title, body)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Post created with title '%s' and body '%s'\n", title, body)
	}
}
