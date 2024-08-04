package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Niflnir/Dreame/internal/models"
	"github.com/Niflnir/Dreame/internal/service"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Niflnir/Dreame/grpc/image"
)

const (
	address = "localhost:50051"
)

type postForCreateUpdate struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	ImageUrl string `json:"image_url"`
}

type postsResponse struct {
	Data       []models.Post `json:"data"`
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
}

type postResponse struct {
	Data       models.Post `json:"data"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
}

type PostControllerImpl struct {
	PostService service.PostService
}

func NewPostControllerImpl(p service.PostService) *PostControllerImpl {
	return &PostControllerImpl{PostService: p}
}

func (c *PostControllerImpl) ListPostHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := c.PostService.ListPosts()
	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		log.Error().Msgf("%v", err)
		return
	}

	res := postsResponse{
		Data:       posts,
		StatusCode: http.StatusOK,
		Message:    "Successfully retrieved posts",
	}

	res.sendJsonResponse(w)
}

func (c *PostControllerImpl) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var p postForCreateUpdate
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusBadRequest)
		log.Error().Msgf("%v", err)
		return
	}

	var postCreated models.Post
	postCreated, err = c.PostService.CreatePost(p.Title, p.Body)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusBadRequest)
		log.Error().Msgf("%v", err)
		return
	}

	var res postResponse
	res = postResponse{
		Data:       postCreated,
		StatusCode: http.StatusCreated,
		Message:    "Successfully created post",
	}

	res.sendJsonResponse(w)
}

func (c *PostControllerImpl) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_num := validateAndExtractIdParameter(w, vars)

	err := c.PostService.DeletePost(id_num)
	if err != nil {
		http.Error(w, "Failed to delete post", http.StatusBadRequest)
		log.Error().Msgf("%v", err)
		return
	}

	res := postResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully deleted post",
	}

	res.sendJsonResponse(w)
}

func (c *PostControllerImpl) UpdatePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_num := validateAndExtractIdParameter(w, vars)

	var u postForCreateUpdate
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedPost models.Post
	updatedPost, err = c.PostService.UpdatePost(id_num, u.Title, u.Body, u.ImageUrl)

	res := postResponse{
		Data:       updatedPost,
		StatusCode: http.StatusOK,
		Message:    "Successfully updated post",
	}

	res.sendJsonResponse(w)
}

func (c *PostControllerImpl) GenerateImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_num := validateAndExtractIdParameter(w, vars)
	post, err := c.PostService.GetPostById(id_num)

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Msgf("Unable to connect to grpc server: %v", err)
	}
	defer conn.Close()

	client := pb.NewImageGeneratorClient(conn)
	req := &pb.ImageRequest{Prompt: post.Body}
	res, err := client.GetImageUrl(context.Background(), req)
	if err != nil {
		log.Error().Msgf("%v", err)
		return
	}

	log.Info().Msgf("Response: %v", res)

	updatedPost, err := c.PostService.UpdatePost(id_num, "", "", res.GetImageUrl())
	if err != nil {
		log.Error().Msgf("%v", err)
		return
	}

	postResponse := postResponse{
		Data:       updatedPost,
		StatusCode: http.StatusOK,
		Message:    "Successfully generated image for post",
	}

	postResponse.sendJsonResponse(w)
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
