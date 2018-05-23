package ml

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// BuildImageFromDockerfile build the image thanks to the Dockerfile
// given when the request is submitted.
func BuildImageFromDockerfile(archivePath string, tags []string) error {
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
	resp, err := cli.ImageBuild(context.Background(), dockerBuildContext, opts)
	if err != nil {
		return fmt.Errorf("could not build the image: %s", err)
	}
	defer resp.Body.Close()
	return nil
}

// RunImage run the image and get the result.
func RunImage() error {
	_, err := client.NewEnvClient()
	if err != nil {
		return fmt.Errorf("could not connect to the docker environnement: %v", err)
	}

	return nil
}
