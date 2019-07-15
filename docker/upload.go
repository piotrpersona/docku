package docker

import (
	"fmt"
	"sync"

	"github.com/docker/docker/client"
	"github.com/piotrpersona/docker-upload/config"
)

func uploadImage(cli client.APIClient, sourceImageName, destinationImageName string, wg *sync.WaitGroup) {
	defer wg.Done()
	sourceImage, err := pull(cli, sourceImageName)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("There was an error while pulling: %s\n", sourceImageName)
	}
	destinationImage := tag(sourceImage, destinationImageName)
	push(destinationImage)
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
