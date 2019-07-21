package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

func tag(cli client.APIClient, sourceImage, destinationImage string) (err error) {
	err = cli.ImageTag(context.Background(), sourceImage, destinationImage)
	if err != nil {
		return
	}
	fmt.Printf("Tagged %s with %s\n", sourceImage, destinationImage)
	return
}
