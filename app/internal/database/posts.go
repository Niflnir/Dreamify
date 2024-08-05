package database

import (
	"database/sql"

	"github.com/Niflnir/Dreame/internal/models"
	"github.com/rs/zerolog/log"
)

type PostRepository interface {
	ListPosts() ([]models.Post, error)
	CreatePost(title string, body string) (models.Post, error)
	DeletePost(id int) error
	UpdatePost(id int, title string, body string, image_url string) (models.Post, error)
	GetPostById(id int) (models.Post, error)
}

type PostRepositoryImpl struct {
	DB *sql.DB
}

func NewPostRepositoryImpl(db *sql.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{DB: db}
}

func (r *PostRepositoryImpl) ListPosts() ([]models.Post, error) {
	rows, err := r.DB.Query("SELECT id, title, body, TO_CHAR(date_created, 'DD-MM-YYYY') as date, image_url FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Id, &post.Title, &post.Body, &post.DateCreated, &post.ImageUrl)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepositoryImpl) CreatePost(title string, body string) (models.Post, error) {
	stmt, err := r.DB.Prepare("INSERT INTO posts(title,body) VALUES($1,$2) RETURNING id, title, body, date_created, image_url")
	if err != nil {
		log.Error().Msgf("Failed to prepare statement: %v", err)
		return models.Post{}, err
	}
	defer stmt.Close()

	var p models.Post
	err = stmt.QueryRow(title, body).Scan(&p.Id, &p.Title, &p.Body, &p.DateCreated, &p.ImageUrl)
	if err != nil {
		log.Error().Msgf("Failed to create post: %v", err)
		return models.Post{}, err
	}

	log.Info().Msgf("Post with title '%s' and body '%s' created successfully!\n", title, body)

	return p, err
}

func (r *PostRepositoryImpl) DeletePost(id int) error {
	stmt, err := r.DB.Prepare("DELETE FROM posts WHERE id = $1")
	if err != nil {
		log.Error().Msgf("Failed to prepare statement: %v", err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		log.Error().Msgf("Failed to delete post: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Error().Msgf("Error retrieving rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		log.Warn().Msgf("No post found with id: %d", id)
	} else {
		log.Info().Msgf("Post with id '%d' deleted successfully!\n", id)
	}

	return nil
}

func (r *PostRepositoryImpl) UpdatePost(id int, title string, body string, image_url string) (models.Post, error) {
	stmt, err := r.DB.Prepare("UPDATE posts SET title=$1, body=$2, image_url=$3 WHERE id=$4 RETURNING id, title, body, date_created, image_url")
	if err != nil {
		log.Error().Msgf("Failed to prepare statement: %v", err)
		return models.Post{}, err
	}
	defer stmt.Close()

	var p models.Post
	err = stmt.QueryRow(title, body, image_url, id).Scan(&p.Id, &p.Title, &p.Body, &p.DateCreated, &p.ImageUrl)
	if err != nil {
		log.Error().Msgf("Failed to update post: %v", err)
		return models.Post{}, err
	}

	log.Info().Msgf("Post with title '%s' and body '%s' updated successfully!\n", title, body)

	return p, err
}

func (r *PostRepositoryImpl) GetPostById(id int) (models.Post, error) {
	row := r.DB.QueryRow("SELECT * from posts where id=$1", id)

	var p models.Post
	var err error

	err = row.Scan(&p.Id, &p.Title, &p.Body, &p.DateCreated, &p.ImageUrl)

	return p, err
}
