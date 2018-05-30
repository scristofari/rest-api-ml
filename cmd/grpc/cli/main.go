package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/scristofari/rest-api-ml/box"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not load the .env file")
	}
	port := os.Getenv("GRPC_PORT")

	conn, err := grpc.Dial(":"+port, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to backend: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	client := box.NewBoxerClient(conn)
	a, err := client.Run(context.Background(), &box.ArtifactRequest{})
	if err != nil {
		log.Fatalf("could not run : %v", err)
	}
	_ = a
}
