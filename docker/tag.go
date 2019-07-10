package docker

import "fmt"

func tag(sourceImage, destinationImage string) string {
	fmt.Printf("Tagged %s with %s\n", sourceImage, destinationImage)
	return destinationImage
}
