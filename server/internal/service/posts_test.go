package service

import (
	"errors"
	"testing"

	"github.com/Niflnir/Dreame/internal/mocks"
	"github.com/Niflnir/Dreame/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestListPostsSuccess(t *testing.T) {
	// Setup
	mockRepo := &mocks.MockPostRepository{}
	service := NewPostServiceImpl(mockRepo)
	expectedPosts := []models.Post{
		{Id: 1, Title: "Test Post", Body: "Test Body"},
		{Id: 2, Title: "Test Post", Body: "Test Body"},
	}
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
	mockRepo := &mocks.MockPostRepository{}
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

func TestCreatePostSuccess(t *testing.T) {
	// Setup
	mockRepo := &mocks.MockPostRepository{}
	service := NewPostServiceImpl(mockRepo)
	testTitle := "Test Title"
	testBody := "Test Body"
	expectedPost := models.Post{Id: 1, Title: testTitle, Body: testBody}
	mockRepo.On("CreatePost", testTitle, testBody).Return(expectedPost, nil)

	// Execute
	post, err := service.CreatePost(testTitle, testBody)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, post, expectedPost)

	// Verify
	mockRepo.AssertExpectations(t)
}

func TestCreatePostFailure(t *testing.T) {
	// Setup
	mockRepo := &mocks.MockPostRepository{}
	service := NewPostServiceImpl(mockRepo)
	testTitle := "Test Title"
	testBody := "Test Body"
	expectedError := errors.New("Test Error")
	mockRepo.On("CreatePost", testTitle, testBody).Return(models.Post{}, expectedError)

	// Execute
	post, err := service.CreatePost(testTitle, testBody)
	assert.Equal(t, models.Post{}, post)

	// Assert
	assert.Error(t, err, expectedError)

	// Verify
	mockRepo.AssertExpectations(t)
}

func TestDeletePostSuccess(t *testing.T) {
	// Setup
	mockRepo := &mocks.MockPostRepository{}
	service := NewPostServiceImpl(mockRepo)
	testId := 1
	mockRepo.On("DeletePost", testId).Return(nil)

	// Execute
	err := service.DeletePost(testId)

	// Assert
	assert.NoError(t, err)

	// Verify
	mockRepo.AssertExpectations(t)
}

func TestDeletePostFailure(t *testing.T) {
	// Setup
	mockRepo := &mocks.MockPostRepository{}
	service := NewPostServiceImpl(mockRepo)
	testId := 1
	expectedError := errors.New("test error")
	mockRepo.On("DeletePost", testId).Return(expectedError)

	// Execute
	err := service.DeletePost(testId)

	// Assert
	assert.Error(t, err, expectedError)

	// Verify
	mockRepo.AssertExpectations(t)
}
