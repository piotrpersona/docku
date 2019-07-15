package app

import (
	"time"

	"github.com/piotrpersona/docker-upload/docker"
)

func Run(configPath string) {
	imagesMetadata := parseConfig(configPath)
	start := time.Now()
	dockercli := docker.Client()
	docker.Upload(dockercli, imagesMetadata)
	report(start)
}
