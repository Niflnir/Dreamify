package api

import (
	"encoding/json"
	"net/http"
	"strconv"

  "github.com/rs/zerolog/log"
	"github.com/Niflnir/Dreame/database"
	"github.com/gorilla/mux"
)

type postForCreate struct {
  Title string `json:"title"`
  Body  string `json:"body"`
}

type Response struct {
	Data interface{} `json:"data"`
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
}

func ListPostHandler(w http.ResponseWriter, r *http.Request) {
  posts := database.ListPosts(database.DBCon);
  res := Response {
    Data: posts,
    StatusCode: http.StatusOK,
    Message: "Successfully retrieved posts",
  }

  sendJsonResponse(w, res)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var c postForCreate
  err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
	}

	database.CreatePost(database.DBCon, c.Title, c.Body)
  
  res := Response {
    Data: "",
    StatusCode: http.StatusCreated,
    Message: "Successfully created post",
  }

  sendJsonResponse(w, res)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

  vars := mux.Vars(r)
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

	database.DeletePost(database.DBCon, id_num)

  res := Response {
    Data: "",
    StatusCode: http.StatusOK,
    Message: "Successfully deleted post",
  }

  sendJsonResponse(w, res)
}

func sendJsonResponse(w http.ResponseWriter, res Response) {
  jsonData, err := json.Marshal(res)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(jsonData)
}
