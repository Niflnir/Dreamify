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

func CreatePost(db *sql.DB, title string, body string) (Post, error) {
  // Start transaction
  tx, err := db.Begin()
  if err != nil {
		log.Error().Err(err)
    return Post{}, err
  }

  var p Post
  {
    stmt, err := tx.Prepare("INSERT INTO posts(title,body) VALUES($1,$2) RETURNING id, title, body, date_created")
    if err != nil {
      log.Error().Err(err)
      return Post{}, err
    }

    defer stmt.Close()

    err = stmt.QueryRow(title, body).Scan(&p.Id, &p.Title, &p.Body, &p.DateCreated)
    if err != nil {
      log.Error().Err(err)
      return Post{}, err
    }
  }

  // Commit transaction
  {
    err := tx.Commit()
    if err != nil {
      log.Error().Err(err)
      return Post{}, err
    }
  }

  log.Info().Msgf("Post with title '%s' and body '%s' created successfully!\n", title, body)

  return p, err
}

func DeletePost(db *sql.DB, id int) (Post, error) {
  // Start transaction
  tx, err := db.Begin()
  if err != nil {
		log.Error().Err(err)
    return Post{}, err
  }

  var p Post
  {
    stmt, err := tx.Prepare("DELETE FROM posts WHERE id = $1 RETURNING id, title, body, date_created")
    if err != nil {
      log.Error().Err(err)
      return Post{}, err
    }

    defer stmt.Close()

    err = stmt.QueryRow(id).Scan(&p.Id, &p.Title, &p.Body, &p.DateCreated)
    if err != nil {
      log.Error().Err(err)
      return Post{}, err
    }
  }

  // Commit transaction
  {
    err := tx.Commit()
    if err != nil {
      log.Error().Err(err)
      return Post{}, err
    }
  }

	log.Info().Msgf("Post with id '%d' deleted successfully!\n", id)

  return p, err
}

func UpdatePost(db *sql.DB, id int, title string, body string) (Post, error) {
  existingPost, err := getPostById(id, db)
  if err != nil {
    log.Error().Err(err)
    return Post{}, err
  }

  // Do not the field if no value is provided
  if title == "" {
    title = existingPost.Title
  }
  if body == "" {
    body = existingPost.Body
  }

  // Start transaction
  tx, err := db.Begin()
  if err != nil {
		log.Error().Err(err)
    return Post{}, err
  }

  var p Post
  {
    stmt, err := tx.Prepare("UPDATE posts SET title=$1, body=$2 WHERE id=$3")
    if err != nil {
      log.Error().Err(err)
      return Post{}, err
    }

    defer stmt.Close()

    err = stmt.QueryRow(title, body, id).Scan(&p.Id, &p.Title, &p.Body, &p.DateCreated)
    if err != nil {
      log.Error().Err(err)
      return Post{}, err
    }
  }

  // Commit transaction
  {
    err := tx.Commit()
    if err != nil {
      log.Error().Err(err)
      return Post{}, err
    }
  }

	log.Info().Msgf("Post with title '%s' and body '%s' updated successfully!\n", title, body)

  return p, err
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

func getPostById(id int, db *sql.DB) (Post, error) {
  row := db.QueryRow("SELECT * from posts where id=$1", id)

  var p Post 
  var err error

  err = row.Scan(&p.Id, &p.Title, &p.Body, &p.DateCreated)

  return p, err
}
