package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Niflnir/Dreame/database"
	"github.com/gorilla/mux"
)

func TestListPostHandler(t *testing.T) {
	// Prep
	initTest()

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/post", nil)

	// Exec
	ListPostHandler(rr, req)

	// Assert
	assertStatusCode(rr.Result().StatusCode, http.StatusOK, t)
}

func TestCreatePostHandler(t *testing.T) {
	// Prep
	initTest()

	postToCreate := postForCreateUpdate{
		Title: "test title",
		Body:  "test body",
	}
	postToCreateJson, _ := json.Marshal(postToCreate)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/post", bytes.NewBuffer(postToCreateJson))

	// Exec
	CreatePostHandler(rr, req)

	// Assert
	assertStatusCode(rr.Result().StatusCode, http.StatusCreated, t)

	createdPost := getPostFromResponse(rr, t)
	if createdPost.Title != postToCreate.Title || createdPost.Body != postToCreate.Body {
		t.Errorf("Incorrect response data returned!\n")
	}

	// Clean up
	deletePost(t, fmt.Sprintf("%v", createdPost.Id))
}

func TestDeletePostHandler(t *testing.T) {
	// Prep
	initTest()
	postToDelete := createPost(t)

	rr := httptest.NewRecorder()
	id := fmt.Sprintf("%v", postToDelete.Id)
	req := httptest.NewRequest("DELETE", "/post/"+id, nil)
	req = mux.SetURLVars(req, map[string]string{"id": id})

	// Exec
	DeletePostHandler(rr, req)

	// Assert
	assertStatusCode(rr.Result().StatusCode, http.StatusOK, t)

	deletedPost := getPostFromResponse(rr, t)
	if deletedPost.Title != postToDelete.Title || deletedPost.Body != postToDelete.Body {
		t.Errorf("Incorrect response data returned!\n")
	}
}

func TestUpdatePostHandler(t *testing.T) {
	// Prep
	initTest()
	post := createPost(t)

	postToUpdate := postForCreateUpdate{
		Title: "updated test title",
		Body:  "updated test body",
	}
	postToUpdateJson, _ := json.Marshal(postToUpdate)

	rr := httptest.NewRecorder()
	id := fmt.Sprintf("%v", post.Id)
	req := httptest.NewRequest("PUT", "/post/"+id, bytes.NewBuffer(postToUpdateJson))
	req = mux.SetURLVars(req, map[string]string{"id": id})

	// Exec
	UpdatePostHandler(rr, req)

	// Assert
	assertStatusCode(rr.Result().StatusCode, http.StatusOK, t)

	updatedPost := getPostFromResponse(rr, t)
	if updatedPost.Title != postToUpdate.Title || updatedPost.Body != postToUpdate.Body {
		t.Errorf("Incorrect response data returned!\n")
	}

	// Clean up
	deletePost(t, id)
}

func assertStatusCode(actualStatusCode int, expectedStatusCode int, t *testing.T) {
	if actualStatusCode != expectedStatusCode {
		t.Errorf("Status code returned %d, did not match expected code %d\n", actualStatusCode, expectedStatusCode)
	}
}

func createPost(t *testing.T) database.Post {
	postToCreate := postForCreateUpdate{
		Title: "test title",
		Body:  "test body",
	}
	postToCreateJson, _ := json.Marshal(postToCreate)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/post", bytes.NewBuffer(postToCreateJson))

	CreatePostHandler(rr, req)
	post := getPostFromResponse(rr, t)

	return post
}

func deletePost(t *testing.T, id string) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/post/"+id, nil)
	req = mux.SetURLVars(req, map[string]string{"id": id})

	DeletePostHandler(rr, req)
}

func getPostFromResponse(rr *httptest.ResponseRecorder, t *testing.T) database.Post {
	var response postResponse
	err := json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	return *response.Data
}

func initTest() {
	database.ConnectToDB()
}
