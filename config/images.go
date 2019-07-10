package config

type Image struct {
	Registry string `json:"registry"`
	Tag      string `json:"tag"`
}

type ImagesSet map[string]Image

type ImagesMetadata struct {
	Registry string    `json:"registry"`
	Images   ImagesSet `json:"images"`
}
