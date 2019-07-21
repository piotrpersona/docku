package docker

import (
	"fmt"
	"sync"

	"github.com/docker/docker/client"
	"github.com/piotrpersona/docker-upload/config"
)

func uploadImage(cli client.APIClient, sourceImage, destinationImage string, wg *sync.WaitGroup) {
	defer wg.Done()
	err := pull(cli, sourceImage)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("There was an error while pulling: %s\n", sourceImage)
	}
	err = tag(cli, sourceImage, destinationImage)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("There was an error while tagging: %s with %s\n", sourceImage, destinationImage)
	}
	err = push(cli, destinationImage)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("There was an error while pushing: %s\n", destinationImage)
	}
}

func Upload(cli client.APIClient, imagesMetadata *config.ImagesMetadata) {
	var wg sync.WaitGroup
	wg.Add(len(imagesMetadata.Images))
	registry := imagesMetadata.Registry
	for imageName, imageMeta := range imagesMetadata.Images {
		sourceImage := ImageURL(imageMeta.Registry, imageName, imageMeta.Tag)
		destinationImage := ImageURL(registry, imageName, imageMeta.Tag)
		go uploadImage(cli, sourceImage, destinationImage, &wg)
	}
	wg.Wait()
}
