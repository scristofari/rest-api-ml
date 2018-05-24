package main

import (
	"log"

	"github.com/scristofari/rest-api-ml/ml/runner"
)

func main() {
	imageName, err := runner.BuildImageFromArtifact("./ml/fixture/archive.tar")
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Build Done")
	containerID, err := runner.RunImage(imageName)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Start Done")
	s, err := runner.GetStateFromContainerID(containerID)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(s.Status)
}
