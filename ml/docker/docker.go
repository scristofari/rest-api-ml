package docker

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// BuildImageFromArtifact build the image thanks to the Dockerfile
// given when the request is submitted.
func BuildImageFromArtifact(archivePath string, tags []string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return fmt.Errorf("could not connect to the docker environnement: %v", err)
	}

	dockerBuildContext, err := os.Open(archivePath)
	if err != nil {
		return fmt.Errorf("could not open the archive: %s", err)
	}
	defer dockerBuildContext.Close()
	opts := types.ImageBuildOptions{
		Remove:  true,
		NoCache: false,
		Tags:    tags,
	}
	if _, err = cli.ImageBuild(context.Background(), dockerBuildContext, opts); err != nil {
		return fmt.Errorf("could not build the image: %s", err)
	}

	return nil
}

// RunImage run the image and get the result.
func RunImage(imageName string) (string, error) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return "", fmt.Errorf("could not connect to the docker environnement: %v", err)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Tty:   true,
	}, nil, nil, "")
	if err != nil {
		return "", fmt.Errorf("could not create the container: %v", err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", fmt.Errorf("could not start the container: %v", err)
	}

	return resp.ID, nil
}

// GetStateFromContainerID returns the status of the container.
func GetStateFromContainerID(containerID string) (*types.ContainerState, error) {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, fmt.Errorf("could not connect to the docker environnement: %v", err)
	}

	info, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return nil, fmt.Errorf("could not retreive info: %v", err)
	}

	return info.State, nil
}
