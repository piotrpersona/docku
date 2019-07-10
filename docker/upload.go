package docker

import (
	"sync"

	"github.com/piotrpersona/docker-upload/config"
)

func uploadImage(sourceImageName, destinationImageName string, wg *sync.WaitGroup) {
	defer wg.Done()
	pullImageMeasured := measureTime(pull)
	pushImageMeasured := measureTime(push)
	sourceImage := pullImageMeasured(sourceImageName)
	destinationImage := tag(sourceImage, destinationImageName)
	pushImageMeasured(destinationImage)
}

func Upload(imagesMetadata *config.ImagesMetadata) {
	var wg sync.WaitGroup
	wg.Add(len(imagesMetadata.Images))
	registry := imagesMetadata.Registry
	for imageName, imageMeta := range imagesMetadata.Images {
		sourceImage := ImageURL(imageMeta.Registry, imageName, imageMeta.Tag)
		destinationImage := ImageURL(registry, imageName, imageMeta.Tag)
		go uploadImage(sourceImage, destinationImage, &wg)
	}
	wg.Wait()
}
