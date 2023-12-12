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
		fmt.Printf("Post with title '%s' and body '%s' created successfully!\n", title, body)
	}
}

func DeletePost(db *sql.DB, id int) {
  fmt.Println(id)
  _, err := db.Exec("DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Post with id '%d' deleted successfully!\n", id)
	}
}
