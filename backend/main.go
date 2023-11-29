package main

import (
	"net/http"

	"github.com/Niflnir/Dreame/api"
	"github.com/Niflnir/Dreame/database"
	_ "github.com/lib/pq"
)

func main() {
  database.ConnectToDB()

	http.HandleFunc("/post", api.CreatePostHandler)
	http.ListenAndServe(":8080", nil)
}

