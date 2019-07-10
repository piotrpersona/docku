package app

import (
	"time"

	"github.com/piotrpersona/docker-upload/docker"
)

func Run() {
	imagesMetadata := parseConfig()
	start := time.Now()
	docker.Upload(imagesMetadata)
	report(start)
}
