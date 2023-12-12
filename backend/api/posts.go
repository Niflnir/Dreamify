package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Niflnir/Dreame/database"
	"github.com/gorilla/mux"
)

type createPost struct {
  Title string `json:"title"`
  Body  string `json:"body"`
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var c createPost
  err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
	}

	database.CreatePost(database.DBCon, c.Title, c.Body)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

  vars := mux.Vars(r)
  id, ok := vars["id"]
  if !ok {
    fmt.Println("id is missing in parameters")
		http.Error(w, "Bad request", http.StatusBadRequest)
  }

  id_num, err := strconv.Atoi(id)
  if err != nil {
    fmt.Println("id parameter provided is not a valid number")
		http.Error(w, "Bad request", http.StatusBadRequest)
  }

	database.DeletePost(database.DBCon, id_num)
}
