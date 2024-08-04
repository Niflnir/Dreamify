package service

import (
	"errors"
	"testing"

	"github.com/Niflnir/Dreame/internal/models"
	"github.com/stretchr/testify/assert"
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

func TestListPostsSuccess(t *testing.T) {
	// Setup
	mockRepo := &MockPostRepository{}
	service := NewPostServiceImpl(mockRepo)
	expectedPosts := []models.Post{{Id: 1, Title: "Test Post"}}
	mockRepo.On("ListPosts").Return(expectedPosts, nil)

	// Execute
	posts, err := service.ListPosts()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, posts, expectedPosts)

	// Verify
	mockRepo.AssertExpectations(t)
}

func TestListPostsFailure(t *testing.T) {
	// Setup
	mockRepo := &MockPostRepository{}
	service := NewPostServiceImpl(mockRepo)
	expectedError := errors.New("Test Error")
	mockRepo.On("ListPosts").Return([]models.Post{}, expectedError)

	// Execute
	posts, err := service.ListPosts()

	// Assert
	assert.Error(t, err, expectedError)
	assert.Equal(t, []models.Post{}, posts)

	// Verify
	mockRepo.AssertExpectations(t)
}
