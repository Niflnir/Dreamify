package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/Niflnir/Dreame/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Niflnir/Dreame/grpc/image"
)

const (
	address = "localhost:50051"
)

func GenerateImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
  id_num := validateAndExtractIdParameter(w, vars)
  post, err := database.GetPostById(id_num)

  var conn *grpc.ClientConn
	conn, err = grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Msgf("Unable to connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pb.NewImageGeneratorClient(conn)
	req := &pb.ImageRequest{Prompt: post.Body}
	res, err := c.GetImageUrl(context.Background(), req)
	if err != nil {
		log.Error().Msgf("%v", err)
		return
	}

	log.Info().Msgf("Response: %v", res)
  
  updatedPost, err := database.UpdatePost(id_num, "", "", res.GetImageUrl())
  if err != nil {
		log.Error().Msgf("%v", err)
    return
  }

  postResponse := postResponse {
    Data: &updatedPost,
    StatusCode: http.StatusOK,
    Message: "Successfully generated image for post",
  }

  postResponse.sendJsonResponse(w)
}
