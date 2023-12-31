package api

import (
	"encoding/json"
	"net/http"
	"strconv"

  "github.com/rs/zerolog/log"
	"github.com/Niflnir/Dreame/database"
	"github.com/gorilla/mux"
)

type postForCreateUpdate struct {
  Title string `json:"title"`
  Body  string `json:"body"`
}

type postsResponse struct {
	Data []database.Post `json:"data"`
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
}

type postResponse struct {
	Data database.Post `json:"data"`
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
}

func ListPostHandler(w http.ResponseWriter, r *http.Request) {
  posts := database.ListPosts();
  res := postsResponse {
    Data: posts,
    StatusCode: http.StatusOK,
    Message: "Successfully retrieved posts",
  }

  res.sendJsonResponse(w)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var c postForCreateUpdate
  err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
	}

  var postCreated database.Post
  postCreated, err = database.CreatePost(c.Title, c.Body)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
  }
  
  res := postResponse {
    Data: postCreated,
    StatusCode: http.StatusCreated,
    Message: "Successfully created post",
  }

  res.sendJsonResponse(w)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id_num := validateAndExtractIdParameter(w, vars)

  postDeleted, err := database.DeletePost(id_num)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
  }

  res := postResponse {
    Data: postDeleted,
    StatusCode: http.StatusOK,
    Message: "Successfully deleted post",
  }

  res.sendJsonResponse(w)
}

func UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id_num := validateAndExtractIdParameter(w, vars)

	var u postForCreateUpdate
  err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
	}

  var updatedPost database.Post
  updatedPost, err = database.UpdatePost(id_num, u.Title, u.Body)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
  }

  res := postResponse {
    Data: updatedPost,
    StatusCode: http.StatusOK,
    Message: "Successfully updated post",
  }

  res.sendJsonResponse(w)
}

func validateAndExtractIdParameter(w http.ResponseWriter, vars map[string]string) int {
  id, ok := vars["id"]
  if !ok {
    log.Info().Msg("id is missing in parameters")
		http.Error(w, "Bad request", http.StatusBadRequest)
  }

  id_num, err := strconv.Atoi(id)
  if err != nil {
    log.Info().Msg("id parameter provided is not a valid number")
		http.Error(w, "Bad request", http.StatusBadRequest)
  }

  return id_num
}

func (res *postResponse) sendJsonResponse(w http.ResponseWriter) {
  jsonData, err := json.Marshal(res)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(res.StatusCode)
  w.Write(jsonData)
}

func (res *postsResponse) sendJsonResponse(w http.ResponseWriter) {
  jsonData, err := json.Marshal(res)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(res.StatusCode)
  w.Write(jsonData)
}
