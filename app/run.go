package app

import (
	"time"

	"github.com/piotrpersona/docku/docker"
)

// Run will parse provided config and launch upload.
func Run(configPath string) {
	imagesMetadata := parseConfig(configPath)
	start := time.Now()
	dockercli := docker.Client()
	docker.Upload(dockercli, imagesMetadata)
	report(start)
}
