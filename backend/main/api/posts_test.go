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

  postToCreate := postForCreateUpdate {
    Title: "test title",
    Body: "test body",
  }
  postToCreateJson, _ := json.Marshal(postToCreate)

  rr := httptest.NewRecorder()
  req := httptest.NewRequest("POST", "/post", bytes.NewBuffer(postToCreateJson))

  // Exec
  CreatePostHandler(rr, req)

  // Assert
  assertStatusCode(rr.Result().StatusCode, http.StatusCreated, t)

  post := getPostFromResponse(rr, t)
  if post.Title != postToCreate.Title || post.Body != postToCreate.Body {
    t.Errorf("Incorrect response data returned!\n")
  }
}

func TestDeletePostHandler(t *testing.T) {
  // Prep
  initTest()
  post := createPost(t)

  rr := httptest.NewRecorder()
  id := fmt.Sprintf("%d", post.Id)
  req := httptest.NewRequest("DELETE", "/post/" + id, nil)
  req = mux.SetURLVars(req, map[string]string{"id": id})

  // Exec
  DeletePostHandler(rr, req)

  // Assert
  assertStatusCode(rr.Result().StatusCode, http.StatusOK, t)
}

func assertStatusCode(actualStatusCode int, expectedStatusCode int, t *testing.T) {
  if actualStatusCode != expectedStatusCode {
    t.Errorf("Status code returned %d, did not match expected code %d\n", actualStatusCode, expectedStatusCode)
  }
}

func createPost(t *testing.T) database.Post {
  postToCreate := postForCreateUpdate {
    Title: "test title",
    Body: "test body",
  }
  postToCreateJson, _ := json.Marshal(postToCreate)

  rr := httptest.NewRecorder()
  req := httptest.NewRequest("POST", "/post", bytes.NewBuffer(postToCreateJson))

  CreatePostHandler(rr, req)
  post := getPostFromResponse(rr, t)

  return post
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
