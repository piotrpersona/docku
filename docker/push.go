package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func push(cli client.APIClient, image string) (imageName string, err error) {
	imageName = image
	fmt.Printf("Pushing image %s\n", image)
	readCloser, err := cli.ImagePush(context.Background(), image, types.ImagePushOptions{})
	if err != nil {
		return
	}
	defer readCloser.Close()
	fmt.Printf("Push done: %s\n", image)
	return
}
