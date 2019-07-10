package docker

import (
	"fmt"
	"math/rand"
	"time"
)

func push(image string) string {
	fmt.Printf("Pushing image %s\n", image)
	sleep := rand.Intn(800)
	time.Sleep(time.Millisecond * time.Duration(sleep))
	fmt.Printf("Push done: %s\n", image)
	return image
}
