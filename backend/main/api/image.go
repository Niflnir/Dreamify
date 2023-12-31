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
		log.Info().Msgf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewImageGeneratorClient(conn)
	req := &pb.ImageRequest{Prompt: post.Body}
	res, err := c.GetImageUrl(context.Background(), req)
	if err != nil {
		log.Error().Err(err)
		return
	}

	log.Info().Msgf("Response: %s", res.GetImageUrl())
}
