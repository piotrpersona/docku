package docker

import (
	"fmt"
	"sync"

	"github.com/docker/docker/client"
	"github.com/piotrpersona/docker-upload/config"
)

type imageUploadLog struct {
	sourceName, targetName string
	stream                 uploadStream
	err                    dockerError
}

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
	uploadLog chan imageUploadLog,
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
	uploadLog <- imageUploadLog{
		sourceName: sourceImage,
		targetName: destinationImage,
		stream:     streamLog,
		err:        dockerErr,
	}
}

func logUpload(uploadLog chan imageUploadLog) {
	for log := range uploadLog {
		fmt.Println(log)
	}
}

func Upload(cli client.APIClient, imagesMetadata *config.ImagesMetadata) {
	var wg sync.WaitGroup
	numberOfTasks := len(imagesMetadata.Images)
	uploadLog := make(chan imageUploadLog, numberOfTasks)
	wg.Add(numberOfTasks)
	registry := imagesMetadata.Registry
	for imageName, imageMeta := range imagesMetadata.Images {
		sourceImage := ImageURL(imageMeta.Registry, imageName, imageMeta.Tag)
		destinationImage := ImageURL(registry, imageName, imageMeta.Tag)
		go uploadImage(
			cli, sourceImage, destinationImage, &wg, uploadLog)
	}
	wg.Wait()
	close(uploadLog)
	logUpload(uploadLog)
}
