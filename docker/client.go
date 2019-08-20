package docker

import (
	"github.com/docker/docker/client"
)

// Client will return docker client object.
func Client() client.APIClient {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.37"))
	if err != nil {
		panic(err)
	}
	return cli
}
