package docker

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var cli *client.Client

func init() {
	// @TODO Refactor, do not let the connection open, defer close it.
	var err error
	cli, err = client.NewEnvClient()
	if err != nil {
		log.Fatalf("could not connect to the docker environnement: %v", err)
	}
}

// BuildImageFromArtifact build the image thanks to the Dockerfile
// given when the request is submitted.
func BuildImageFromArtifact(archivePath string) (string, error) {
	dockerBuildContext, err := os.Open(archivePath)
	if err != nil {
		return "", fmt.Errorf("could not open the archive: %s", err)
	}
	defer dockerBuildContext.Close()

	imageName := getMD5Hash(dockerBuildContext)

	opts := types.ImageBuildOptions{
		Remove:  true,
		NoCache: false,
		Tags:    []string{imageName},
	}
	if _, err = cli.ImageBuild(context.Background(), dockerBuildContext, opts); err != nil {
		return "", fmt.Errorf("could not build the image: %s", err)
	}

	return imageName, nil
}

// RunImage run the image and get the result.
func RunImage(imageName string) (string, error) {
	ctx := context.Background()
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
	info, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return nil, fmt.Errorf("could not retreive info: %v", err)
	}

	return info.State, nil
}

// StopContainerFromID stop the container if running.
func StopContainerFromID(containerID string) error {
	ctx := context.Background()
	timeout := time.Second * 2
	err := cli.ContainerStop(ctx, containerID, &timeout)
	if err != nil {
		return fmt.Errorf("could not retreive info: %v", err)
	}

	return nil
}

func getMD5Hash(f io.Reader) string {
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(h.Sum(nil))
}
