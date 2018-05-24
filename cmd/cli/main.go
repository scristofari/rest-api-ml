package main

import (
	"log"

	"github.com/scristofari/rest-api-ml/ml/docker"
)

func main() {
	imageName := "api-ml-test2"
	err := docker.BuildImageFromArtifact("./ml/fixture/archive.tar", []string{imageName})
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Build Done")
	containerID, err := docker.RunImage(imageName)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Start Done")
	s, err := docker.GetStateFromContainerID(containerID)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println(s.Status)
}
