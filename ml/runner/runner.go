package runner

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

// BuildImageFromArtifact build the image thanks to the Dockerfile
// given when the request is submitted.
func BuildImageFromArtifact(archivePath string) (string, error) {
	dockerBuildContext, err := os.Open(archivePath)
	if err != nil {
		return "", fmt.Errorf("could not open the archive: %s", err)
	}
	defer dockerBuildContext.Close()

	// the copy will corrupt the tar file apparently.
	// fake it a the moment.
	// imageName, err := getMD5Hash(dockerBuildContext)
	// if err != nil {
	//	return "", err
	// }

	imageName := "api-artifact"
	opts := types.ImageBuildOptions{
		Remove:  true,
		NoCache: false,
		Tags:    []string{imageName},
	}

	cli := getClient()
	defer cli.Close()
	if _, err = cli.ImageBuild(context.Background(), dockerBuildContext, opts); err != nil {
		return "", fmt.Errorf("could not build the image: %s", err)
	}

	return imageName, nil
}

// RunImage run the image and get the result.
func RunImage(imageName string) (string, error) {
	cli := getClient()
	defer cli.Close()
	resp, err := cli.ContainerCreate(context.Background(), &container.Config{
		Image: imageName,
		Tty:   true,
	}, nil, nil, "")
	if err != nil {
		return "", fmt.Errorf("could not create the container: %v", err)
	}
	if err := cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", fmt.Errorf("could not start the container: %v", err)
	}

	return resp.ID, nil
}

// GetStateFromContainerID returns the status of the container.
func GetStateFromContainerID(containerID string) (*types.ContainerState, error) {
	cli := getClient()
	defer cli.Close()
	info, err := cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve info: %v", err)
	}

	return info.State, nil
}

// StopContainerFromID stop the container if running.
func StopContainerFromID(containerID string) error {
	cli := getClient()
	defer cli.Close()
	timeout := time.Second * 2
	err := cli.ContainerStop(context.Background(), containerID, &timeout)
	if err != nil {
		return fmt.Errorf("could not retreive info: %v", err)
	}

	return nil
}

// LogsFromContainerID get the container's STDOUT.
func LogsFromContainerID(containerID string) (io.ReadCloser, error) {
	cli := getClient()
	defer cli.Close()
	return cli.ContainerLogs(context.Background(), containerID, types.ContainerLogsOptions{
		ShowStdout: true,
		Timestamps: false,
	})
}

func getClient() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("could not connect to the docker environnement: %v", err)
	}

	return cli
}

func getMD5Hash(f io.Reader) (string, error) {
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", fmt.Errorf("could not copy: %s", err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
