package main

import (
	"testing"

	"github.com/scristofari/rest-api-ml/ml/runner"
)

func TestCli(t *testing.T) {
	imageName, err := runner.BuildImageFromArtifact("../../ml/fixture/archive.tar")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	containerID, err := runner.RunImage(imageName)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	_, err = runner.GetStateFromContainerID(containerID)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	err = runner.StopContainerFromID(containerID)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
