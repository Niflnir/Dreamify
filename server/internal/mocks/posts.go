package mocks

import (
	"github.com/Niflnir/Dreame/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockPostRepository struct {
	mock.Mock
}

func (m *MockPostRepository) ListPosts() ([]models.Post, error) {
	args := m.Called()
	return args.Get(0).([]models.Post), args.Error(1)
}

func (m *MockPostRepository) CreatePost(title string, body string) (models.Post, error) {
	args := m.Called(title, body)
	return args.Get(0).(models.Post), args.Error(1)
}

func (m *MockPostRepository) DeletePost(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockPostRepository) UpdatePost(id int, title string, body string, image_url string) (models.Post, error) {
	args := m.Called(id, title, body, image_url)
	return args.Get(0).(models.Post), args.Error(1)
}

func (m *MockPostRepository) GetPostById(id int) (models.Post, error) {
	args := m.Called(id)
	return args.Get(0).(models.Post), args.Error(1)
}
