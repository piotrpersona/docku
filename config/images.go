package config

type image struct {
	Registry string `json,yaml:"registry"`
	Tag      string `json,yaml:"tag"`
}

type imagesSet map[string]image

// ImagesMetadata presents list of images that will be uploaded to registry.
type ImagesMetadata struct {
	Registry string    `json,yaml:"registry"`
	Images   imagesSet `json,yaml:"images"`
}
