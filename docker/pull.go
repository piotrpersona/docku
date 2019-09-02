package docker

import (
	"github.com/piotrpersona/docku/config"
	"context"
	"fmt"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

)

func pull(cli client.APIClient, image config.ImageURL) (stream []byte, err error) {
	fmt.Printf("Pulling image %s\n", image)
	streamReader, err := cli.ImagePull(context.Background(), string(image),
		types.ImagePullOptions{})
	if err != nil {
		return
	}
	defer streamReader.Close()
	stream, err = ioutil.ReadAll(streamReader)
	if err != nil {
		return
	}
	fmt.Printf("Pull done: %s\n", image)
	return
}
