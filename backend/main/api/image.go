package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Niflnir/Dreame/grpc/image"
)

const (
	address = "localhost:50051"
)

func GenerateImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_ = validateAndExtractIdParameter(w, vars)

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Info().Msgf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewImageGeneratorClient(conn)

	req := &pb.ImageRequest{Prompt: "A cute brown corgi."}
	res, err := c.GetImageUrl(context.Background(), req)
	if err != nil {
		log.Error().Err(err)
		return
	}

	log.Printf("Response: %s", res.GetImageUrl())
}
