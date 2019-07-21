package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func pull(cli client.APIClient, image string) (err error) {
	fmt.Printf("Pulling image %s\n", image)
	readCloser, err := cli.ImagePull(context.Background(), image, types.ImagePullOptions{})
	if err != nil {
		return
	}
	defer readCloser.Close()
	fmt.Printf("Pull done: %s\n", image)
	return
}
