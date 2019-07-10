package app

import (
	"time"

	"github.com/piotrpersona/docker-upload/docker"
)

func Run(configPath string) {
	imagesMetadata := parseConfig(configPath)
	start := time.Now()
	docker.Upload(imagesMetadata)
	report(start)
}
