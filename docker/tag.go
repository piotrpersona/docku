package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

func tag(cli client.APIClient, sourceImage, destinationImage string) (destinationImageName string, err error) {
	err = cli.ImageTag(context.Background(), sourceImage, destinationImage)
	if err != nil {
		return "", err
	}
	fmt.Printf("Tagged %s with %s\n", sourceImage, destinationImage)
	destinationImageName = destinationImage
	return destinationImage, nil
}
