package main

import (
	"net/http"

	"github.com/Niflnir/Dreame/internal/database"
	"github.com/Niflnir/Dreame/internal/handlers"
	"github.com/Niflnir/Dreame/internal/service"
	"github.com/gorilla/mux"
)

func main() {
	db := database.ConnectToDB()
	postRepo := database.NewPostRepositoryImpl(db)
	postService := service.NewPostServiceImpl(postRepo)
	postController := handlers.NewPostControllerImpl(postService)
	imageController := handlers.NewImageControllerImpl(postService)

	r := mux.NewRouter()

	// Post
	r.HandleFunc("/posts", postController.ListPostHandler).Methods("GET")
	r.HandleFunc("/posts", postController.CreatePostHandler).Methods("POST")
	r.HandleFunc("/posts/{id}", postController.DeletePostHandler).Methods("DELETE")
	r.HandleFunc("/posts/{id}", postController.UpdatePostHandler).Methods("PUT")
	r.HandleFunc("/generate-image", imageController.GenerateImageHandler).Methods("POST")

	http.ListenAndServe(":8080", r)
}
