package main

import (
	"net/http"

	"github.com/Niflnir/Dreame/api"
	"github.com/Niflnir/Dreame/database"
	"github.com/gorilla/mux"
)

func main() {
  database.ConnectToDB()
  r := mux.NewRouter()

	r.HandleFunc("/post", api.CreatePostHandler)
  r.HandleFunc("/post/{id}", api.DeletePostHandler)
	http.ListenAndServe(":8080", r)
}

