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
	var c postForCreateUpdate
  err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
	}

	database.CreatePost(database.DBCon, c.Title, c.Body)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
  }
  
  res := Response {
    Data: "",
    StatusCode: http.StatusCreated,
    Message: "Successfully created post",
  }

  sendJsonResponse(w, res)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id_num := validateAndExtractIdParameter(w, vars)

  p, err := database.DeletePost(database.DBCon, id_num)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
  }

  res := Response {
    Data: p,
    StatusCode: http.StatusOK,
    Message: "Successfully deleted post",
  }

  sendJsonResponse(w, res)
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

  var p database.Post
  p, err = database.UpdatePost(database.DBCon, id_num, u.Title, u.Body)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
  }

  res := Response {
    Data: p,
    StatusCode: http.StatusOK,
    Message: "Successfully updated post",
  }

  sendJsonResponse(w, res)
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

func sendJsonResponse(w http.ResponseWriter, res Response) {
  jsonData, err := json.Marshal(res)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(jsonData)
}
