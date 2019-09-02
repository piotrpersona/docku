package docker

import (
	"context"
	"fmt"

	"github.com/piotrpersona/docku/config"

	"github.com/docker/docker/client"
)

func tag(cli client.APIClient, sourceImage, destinationImage config.ImageURL) (err error) {
	err = cli.ImageTag(context.Background(), string(sourceImage), string(destinationImage))
	if err != nil {
		return
	}
	fmt.Printf("Tagged %s with %s\n", sourceImage, destinationImage)
	return
}
