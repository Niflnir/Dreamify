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
  ImageUrl string
}

func ListPosts() []Post {
  rows, err := DBCon.Query("SELECT id, title, body, TO_CHAR(date_created, 'DD-MM-YYYY') as date, image_url FROM posts")
	if err != nil {
    log.Error().Msgf("%v", err)
  }
  defer rows.Close()

  var posts []Post
  for rows.Next() {
    post := rowsToPost(rows)
    posts = append(posts, post)
  }

  return posts
}

func CreatePost(title string, body string) (Post, error) {
  // Start transaction
  tx, err := DBCon.Begin()
  if err != nil {
    log.Error().Msgf("%v", err)
    return Post{}, err
  }

  var p Post
  {
    stmt, err := tx.Prepare("INSERT INTO posts(title,body) VALUES($1,$2) RETURNING id, title, body, date_created, image_url")
    if err != nil {
      log.Error().Msgf("%v", err)
      return Post{}, err
    }

    defer stmt.Close()

    err = stmt.QueryRow(title, body).Scan(&p.Id, &p.Title, &p.Body, &p.DateCreated, &p.ImageUrl)
    if err != nil {
      log.Error().Msgf("%v", err)
      return Post{}, err
    }
  }

  // Commit transaction
  {
    err := tx.Commit()
    if err != nil {
      log.Error().Msgf("%v", err)
      return Post{}, err
    }
  }

  log.Info().Msgf("Post with title '%s' and body '%s' created successfully!\n", title, body)

  return p, err
}

func DeletePost(id int) (Post, error) {
  // Start transaction
  tx, err := DBCon.Begin()
  if err != nil {
    log.Error().Msgf("%v", err)
    return Post{}, err
  }

  var p Post
  {
    stmt, err := tx.Prepare("DELETE FROM posts WHERE id = $1 RETURNING id, title, body, date_created, image_url")
    if err != nil {
      log.Error().Msgf("%v", err)
      return Post{}, err
    }

    defer stmt.Close()

    err = stmt.QueryRow(id).Scan(&p.Id, &p.Title, &p.Body, &p.DateCreated, &p.ImageUrl)
    if err != nil {
      log.Error().Msgf("%v", err)
      return Post{}, err
    }
  }

  // Commit transaction
  {
    err := tx.Commit()
    if err != nil {
      log.Error().Msgf("%v", err)
      return Post{}, err
    }
  }

	log.Info().Msgf("Post with id '%d' deleted successfully!\n", id)

  return p, err
}

func UpdatePost(id int, title string, body string, image_url string) (Post, error) {
  existingPost, err := GetPostById(id)
  log.Info().Msgf("%v", existingPost)
  if err != nil {
    log.Error().Msgf("%v", err)
    return Post{}, err
  }

  // Do not update the field if no value is provided
  if title == "" {
    title = existingPost.Title
  }
  if body == "" {
    body = existingPost.Body
  }
  if image_url == "" {
    image_url = existingPost.ImageUrl
  }

  // Start transaction
  tx, err := DBCon.Begin()
  if err != nil {
    log.Error().Msgf("%v", err)
    return Post{}, err
  }

  var p Post
  {
    stmt, err := tx.Prepare("UPDATE posts SET title=$1, body=$2, image_url=$3 WHERE id=$4 RETURNING id, title, body, date_created, image_url")
    if err != nil {
      log.Error().Msgf("%v", err)
      return Post{}, err
    }

    defer stmt.Close()

    log.Info().Msgf("title: %v, body: %v, image_url: %v, id: %v",title, body, image_url, id)
    err = stmt.QueryRow(title, body, image_url, id).Scan(&p.Id, &p.Title, &p.Body, &p.DateCreated, &p.ImageUrl)
    if err != nil {
      log.Error().Msgf("%v", err)
      return Post{}, err
    }
  }

  // Commit transaction
  {
    err := tx.Commit()
    if err != nil {
      log.Error().Msgf("%v", err)
      return Post{}, err
    }
  }

	log.Info().Msgf("Post with title '%s' and body '%s' updated successfully!\n", title, body)

  return p, err
}

func rowsToPost(rows *sql.Rows) Post {
  var post Post 
  var err error

  err = rows.Scan(&post.Id, &post.Title, &post.Body, &post.DateCreated, &post.ImageUrl)
  if err != nil {
    log.Error().Msgf("%v", err)
  }

  return post
}

func GetPostById(id int) (Post, error) {
  row := DBCon.QueryRow("SELECT * from posts where id=$1", id)

  var p Post 
  var err error

  err = row.Scan(&p.Id, &p.Title, &p.Body, &p.DateCreated, &p.ImageUrl)

  return p, err
}
