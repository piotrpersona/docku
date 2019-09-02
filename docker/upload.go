package docker

import (
	"fmt"
	"os"
	"sync"

	"github.com/docker/docker/client"
	"github.com/piotrpersona/docku/config"
)

type imageUploadLog struct {
	sourceName, targetName config.ImageURL
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
	destinationImage config.ImageURL,
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

// Upload will upload docker images defined in metadata to remote registry.
func Upload(cli client.APIClient, imagesMetadata *config.ImagesMetadata) {
	var wg sync.WaitGroup
	numberOfTasks := len(imagesMetadata.Images)
	uploadLog := make(chan imageUploadLog, numberOfTasks)
	wg.Add(numberOfTasks)
	for _, image := range imagesMetadata.Images {
		err := image.Validate()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sourceImage := image
		destinationImage := image.ReplaceRegistry(imagesMetadata.Registry)
		go uploadImage(
			cli, sourceImage, destinationImage, &wg, uploadLog)
	}
	wg.Wait()
	close(uploadLog)
	logUpload(uploadLog)
}
