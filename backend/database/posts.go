package database

import (
	"database/sql"

  "github.com/rs/zerolog/log"
)

type Post struct {
  Id int 
  Title string
  Body  string
  DateCreated string
}

func ListPosts(db *sql.DB) []Post {
  rows, err := db.Query("SELECT id, title, body, TO_CHAR(date_created, 'DD-MM-YYYY') as date FROM posts")
	if err != nil {
		log.Error().Err(err)
  }
  defer rows.Close()

  var posts []Post
  for rows.Next() {
    var post Post 

    err = rows.Scan(&post.Id, &post.Body, &post.Title, &post.DateCreated)
    if err != nil {
      log.Error().Err(err)
    }

    posts = append(posts, post)
  }

  return posts
}

func CreatePost(db *sql.DB, title string, body string) {
	_, err := db.Exec("INSERT INTO posts(title,body) VALUES($1,$2)", title, body)
	if err != nil {
		log.Error().Err(err)
	} else {
		log.Info().Msgf("Post with title '%s' and body '%s' created successfully!\n", title, body)
	}
}

func DeletePost(db *sql.DB, id int) {
  _, err := db.Exec("DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		log.Error().Err(err)
	} else {
		log.Info().Msgf("Post with id '%d' deleted successfully!\n", id)
	}
}
