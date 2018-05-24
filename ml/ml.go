package ml

import (
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/scristofari/rest-api-ml/ml/docker"
)

type Runner interface {
	BuildImageFromArtifact(archivePath string) (string, error)
	RunImage(imageName string) (string, error)
	GetStateFromContainerID(containerID string) (*types.ContainerState, error)
	StopContainerFromID(containerID string) error
}

type Storage interface {
	AddNewArtifact(artifact *ArtifactInfo) error
	AddNewEventForArtifact(event map[string]interface{}, uuid int64) error
	GetArtifactInfoFromUUID(uuid int64) (*ArtifactInfo, error)
}

type ArtifactInfo struct {
	UUID   int64                  `json:"uuid"`   // @TODO lib UUID
	Status string                 `json:"status"` // @TODO Enum or {event}_{status} default 'pending'
	Events map[string]interface{} `json:"events"` // @TODO map case with the proper field
	Result struct {
		TestLoss     float64 `json:"test_loss"`
		TestAccuracy float64 `json:"test_accuracy"`
	} `json:"result"`
}

// LaunchArtifact will build and run the project.
// @TODO Store each event -> interface Storage
func LaunchArtifact(artifactPath string) error {
	imageName, err := docker.BuildImageFromArtifact(artifactPath)
	if err != nil {
		return fmt.Errorf("could not build the image %v", err)
	}
	containerID, err := docker.RunImage(imageName)
	if err != nil {
		return fmt.Errorf("could not run the image %v", err)
	}
	_ = containerID
	return nil
}

// GetArtifactInfo will retrieve info on the launch's progress
func GetArtifactInfo(uuid int64) (*ArtifactInfo, error) {
	return nil, fmt.Errorf("Not implemented")
}

// GetArtifactLogs will retrieve logs from a specific event
func GetArtifactLogs(artifactID string, event string) (*io.Reader, error) {
	return nil, fmt.Errorf("Not implemented")
}
