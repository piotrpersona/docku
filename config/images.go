package config

type image struct {
	Registry string `json:"registry"`
	Tag      string `json:"tag"`
}

type imagesSet map[string]image

// ImagesMetadata presents list of images that will be uploaded to registry.
type ImagesMetadata struct {
	Registry string    `json:"registry"`
	Images   imagesSet `json:"images"`
}
