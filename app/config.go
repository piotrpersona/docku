package app

import (
	"fmt"
	"os"

	"github.com/piotrpersona/docker-upload/config"
)

func parseConfig(configPath string) *config.ImagesMetadata {
	imagesMetadata, err := config.Read(configPath)
	if err != nil {
		fmt.Printf("There was an error while reading: %s", configPath)
		os.Exit(1)
	}
	return imagesMetadata
}
