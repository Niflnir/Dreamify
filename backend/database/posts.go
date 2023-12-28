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
    post := rowsToPost(rows)
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

func DeletePost(db *sql.DB, id int) (Post, error) {
  postToDelete := getPostById(id, db)

  _, err := db.Exec("DELETE FROM posts WHERE id = $1", id)
	if err != nil {
		log.Error().Err(err)
	} else {
		log.Info().Msgf("Post with id '%d' deleted successfully!\n", id)
	}

  return postToDelete, err
}

func UpdatePost(db *sql.DB, id int, title string, body string) (Post, error) {
  existingPost := getPostById(id, db)

  if title == "" {
    title = existingPost.Title
  }

  if body == "" {
    body = existingPost.Body
  }

  _, err := db.Exec("UPDATE posts SET title=$1, body=$2 WHERE id=$3", title, body, id)
	if err != nil {
		log.Error().Err(err)
	} else {
		log.Info().Msgf("Post with title '%s' and body '%s' updated successfully!\n", title, body)
	}

  updatedPost := getPostById(id, db)
  return updatedPost, err
}

func rowsToPost(rows *sql.Rows) Post {
  var post Post 
  var err error

  err = rows.Scan(&post.Id, &post.Title, &post.Body, &post.DateCreated)
  if err != nil {
    log.Error().Err(err)
  }

  return post
}

func getPostById(id int, db *sql.DB) Post {
  row := db.QueryRow("SELECT * from posts where id=$1", id)

  var post Post 
  var err error

  err = row.Scan(&post.Id, &post.Title, &post.Body, &post.DateCreated)
  if err != nil {
    log.Error().Err(err)
  }

  return post
}
