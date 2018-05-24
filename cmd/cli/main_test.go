package main

import (
	"testing"

	"github.com/scristofari/rest-api-ml/ml/docker"
)

func TestCli(t *testing.T) {
	imageName, err := docker.BuildImageFromArtifact("../../ml/fixture/archive.tar")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	containerID, err := docker.RunImage(imageName)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	_, err = docker.GetStateFromContainerID(containerID)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	err = docker.StopContainerFromID(containerID)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
