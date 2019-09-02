package config

import (
	"fmt"
	"net/url"
	"strings"
)

// ImageURL represents docker image URL.
type ImageURL string

// Validate will check if ImageURL is in a valid format.
func (imageURL ImageURL) Validate() (err error) {
	imageURLSplit := strings.Split(string(imageURL), ":")
	if len(imageURLSplit) > 2 {
		return fmt.Errorf("Invalid image URL:%s", imageURL)
	}
	imagePrefix := imageURLSplit[0]
	_, err = url.Parse(imagePrefix)
	return err
}

func buildURL(registry, imageName string) string {
	registryTrimmed := strings.Trim(registry, "/")
	imageTrimmed := strings.Trim(imageName, "/")
	return registryTrimmed + "/" + imageTrimmed
}

// ReplaceRegistry will return ImageURL with provided registry.
func (imageURL ImageURL) ReplaceRegistry(registry string) (url ImageURL) {
	sourceImageURL := string(imageURL)
	containsRegistry := strings.Count(sourceImageURL, "/") != 0
	if !containsRegistry {
		url = ImageURL(buildURL(registry, sourceImageURL))
		return
	}
	urlSplitted := strings.SplitN(sourceImageURL, "/", 2)
	imageName := urlSplitted[1]
	url = ImageURL(buildURL(registry, imageName))
	return
}

// ImagesMetadata presents list of images that will be uploaded to registry.
type ImagesMetadata struct {
	Registry string     `json,yaml:"registry"`
	Images   []ImageURL `json,yaml:"images"`
}
