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

  r.HandleFunc("/post", api.ListPostHandler).Methods("GET")
	r.HandleFunc("/post", api.CreatePostHandler).Methods("POST")
  r.HandleFunc("/post/{id}", api.DeletePostHandler).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}

