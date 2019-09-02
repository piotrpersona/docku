package docker

import (
	"github.com/piotrpersona/docku/config"
	"context"
	"fmt"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func push(cli client.APIClient, image config.ImageURL) (stream []byte, err error) {
	fmt.Printf("Pushing image %s\n", image)
	streamReader, err := cli.ImagePush(context.Background(), string(image),
		types.ImagePushOptions{
			All:          true,
			RegistryAuth: "123",
		})
	if err != nil {
		return
	}
	defer streamReader.Close()
	stream, err = ioutil.ReadAll(streamReader)
	if err != nil {
		return
	}
	fmt.Printf("Push done: %s\n", image)
	return
}
