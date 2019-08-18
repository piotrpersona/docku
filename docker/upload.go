package docker

import (
	"fmt"
	"sync"

	"github.com/docker/docker/client"
	"github.com/piotrpersona/docker-upload/config"
)

type dockerError struct {
	push, tag, pull string
}

type uploadStream struct {
	pull, push string
}

func uploadImage(
	cli client.APIClient, sourceImage,
	destinationImage string,
	wg *sync.WaitGroup,
	streamBuffer chan uploadStream,
	errorBuffer chan dockerError,
) {
	defer wg.Done()
	dockerErr := dockerError{}
	pullStream, pullErr := pull(cli, sourceImage)
	if pullErr != nil {
		dockerErr.pull = pullErr.Error()
	}
	tagErr := tag(cli, sourceImage, destinationImage)
	if tagErr != nil {
		dockerErr.tag = tagErr.Error()
	}
	pushStream, pushErr := push(cli, destinationImage)
	if pushErr != nil {
		dockerErr.push = pushErr.Error()
	}
	streamLog := uploadStream{
		pull: string(pullStream),
		push: string(pushStream),
	}
	streamBuffer <- streamLog
	errorBuffer <- dockerErr
}

func logUpload(streamBuffer chan uploadStream, errorBuffer chan dockerError) {
	for err := range errorBuffer {
		fmt.Println(err)
	}
	for stream := range streamBuffer {
		fmt.Println(stream)
	}
}

func Upload(cli client.APIClient, imagesMetadata *config.ImagesMetadata) {
	var wg sync.WaitGroup
	numberOfTasks := len(imagesMetadata.Images)
	streamBuffer := make(chan uploadStream, numberOfTasks)
	errorBuffer := make(chan dockerError, numberOfTasks)
	wg.Add(numberOfTasks)
	registry := imagesMetadata.Registry
	for imageName, imageMeta := range imagesMetadata.Images {
		sourceImage := ImageURL(imageMeta.Registry, imageName, imageMeta.Tag)
		destinationImage := ImageURL(registry, imageName, imageMeta.Tag)
		go uploadImage(
			cli, sourceImage, destinationImage, &wg, streamBuffer, errorBuffer)
	}
	wg.Wait()
	close(streamBuffer)
	close(errorBuffer)
	logUpload(streamBuffer, errorBuffer)
}
