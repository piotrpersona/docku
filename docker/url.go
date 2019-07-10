package docker

import (
	"fmt"
	"strings"
)

func ImageURL(registry, name, tag string) string {
	validRegistry := strings.TrimRight(registry, "/")
	validName := strings.TrimRight(strings.TrimLeft(name, "/"), ":")
	validTag := strings.TrimLeft(tag, ":")
	imageURL := validRegistry + "/" + validName + ":" + validTag
	fmt.Println(imageURL)
	return imageURL
}
