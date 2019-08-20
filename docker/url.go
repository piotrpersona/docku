package docker

import (
	"fmt"
	"strings"
)

// ImageURL will trim redundant character from Image URL.
func ImageURL(registry, name, tag string) string {
	validRegistry := strings.TrimRight(registry, "/")
	validName := strings.TrimRight(strings.TrimLeft(name, "/"), ":")
	validTag := strings.TrimLeft(tag, ":")
	imageURL := validRegistry + "/" + validName + ":" + validTag
	fmt.Println(imageURL)
	return imageURL
}
