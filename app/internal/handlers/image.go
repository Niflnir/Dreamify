package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Niflnir/Dreame/internal/service"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Niflnir/Dreame/grpc/image"
)

const (
	address = "localhost:50051"
)

type GenerateImageRequest struct {
	PostId int `json:"postId"`
}

type ImageControllerImpl struct {
	PostService service.PostService
}

func NewImageControllerImpl(p service.PostService) *ImageControllerImpl {
	return &ImageControllerImpl{PostService: p}
}

func (c *ImageControllerImpl) GenerateImageHandler(w http.ResponseWriter, r *http.Request) {
	var g GenerateImageRequest
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		log.Error().Msgf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	postId := g.PostId
	post, err := c.PostService.GetPostById(postId)
	if err != nil {
		log.Error().Msgf("Failed to retrieve post: %v", err)
		http.Error(w, "Failed to generate image", http.StatusInternalServerError)
		return
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Msgf("Unable to connect to grpc server: %v", err)
		http.Error(w, "Failed to generate image", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := pb.NewImageGeneratorClient(conn)
	req := &pb.ImageRequest{Prompt: post.Body}
	res, err := client.GetImageUrl(context.Background(), req)
	if err != nil {
		log.Error().Msgf("Failed to generate image via python client %v", err)
		http.Error(w, "Failed to generate image", http.StatusInternalServerError)
		return
	}

	updatedPost, err := c.PostService.UpdatePost(postId, post.Title, post.Body, res.GetImageUrl())
	if err != nil {
		log.Error().Msgf("Failed to update post with image url: %v", err)
		http.Error(w, "Failed to generate image", http.StatusInternalServerError)
		return
	}

	postResponse := postResponse{
		Data:       updatedPost,
		StatusCode: http.StatusOK,
		Message:    "Successfully generated image for post",
	}

	postResponse.sendJsonResponse(w)
}
