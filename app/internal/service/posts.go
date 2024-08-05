package service

import (
	db "github.com/Niflnir/Dreame/internal/database"
	"github.com/Niflnir/Dreame/internal/models"
)

type PostService interface {
	ListPosts() ([]models.Post, error)
	CreatePost(title string, body string) (models.Post, error)
	DeletePost(id int) error
	UpdatePost(id int, title string, body string, image_url string) (models.Post, error)
	GetPostById(id int) (models.Post, error)
}

type PostServiceImpl struct {
	Repo db.PostRepository
}

func NewPostServiceImpl(repo db.PostRepository) *PostServiceImpl {
	return &PostServiceImpl{Repo: repo}
}

func (s *PostServiceImpl) ListPosts() ([]models.Post, error) {
	return s.Repo.ListPosts()
}

func (s *PostServiceImpl) CreatePost(title string, body string) (models.Post, error) {
	return s.Repo.CreatePost(title, body)
}

func (s *PostServiceImpl) DeletePost(id int) error {
	return s.Repo.DeletePost(id)
}

func (s *PostServiceImpl) UpdatePost(id int, title string, body string, image_url string) (models.Post, error) {
	return s.Repo.UpdatePost(id, title, body, image_url)
}

func (s *PostServiceImpl) GetPostById(id int) (models.Post, error) {
	return s.Repo.GetPostById(id)
}
