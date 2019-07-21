package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func push(cli client.APIClient, image string) (err error) {
	fmt.Printf("Pushing image %s\n", image)
	readCloser, err := cli.ImagePush(context.Background(), image, types.ImagePushOptions{
		All: true,
	})
	if err != nil {
		return
	}
	defer readCloser.Close()
	fmt.Printf("Push done: %s\n", image)
	return
}
