package ml

import (
	"fmt"

	"github.com/scristofari/rest-api-ml/ml/docker"
)

type Status struct {
	ArtifactPath string
	ImageName    string
	ContainerID  string
}

func LaunchArtifact(artifactPath string) (*Status, error) {
	imageName, err := docker.BuildImageFromArtifact(artifactPath)
	if err != nil {
		return nil, fmt.Errorf("could not build the image %v", err)
	}
	containerID, err := docker.RunImage(imageName)
	if err != nil {
		return nil, fmt.Errorf("could not run the image %v", err)
	}

	return &Status{
		ArtifactPath: artifactPath,
		ImageName:    imageName,
		ContainerID:  containerID,
	}, nil
}
