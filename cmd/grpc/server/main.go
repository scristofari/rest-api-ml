package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/scristofari/rest-api-ml/box"
	"google.golang.org/grpc"
)

type BoxServer struct{}

func (BoxServer) Run(context.Context, *box.ArtifactRequest) (*box.Artifact, error) {
	return nil, fmt.Errorf("Not implemented")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not load the .env file")
	}

	s := grpc.NewServer()
	box.RegisterBoxerServer(s, BoxServer{})

	port := os.Getenv("GRPC_PORT")
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("could not listen to :%s: %v", port, err)
	}
	log.Fatal(s.Serve(l))
}
