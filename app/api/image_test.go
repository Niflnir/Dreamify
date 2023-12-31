package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGenerateImageHandler(t *testing.T) {
  // Prep
  initTest()
  postFixture := createPost(t)

  rr := httptest.NewRecorder()
  id := fmt.Sprintf("%v", postFixture.Id)
  req := httptest.NewRequest("POST", "/post/generate-image" + id, nil)
  req = mux.SetURLVars(req, map[string]string{"id": id})
  
  // Exec
  GenerateImageHandler(rr, req)

  // Assert
  assertStatusCode(rr.Result().StatusCode, http.StatusOK, t)
}

