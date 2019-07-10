package docker

import (
	"fmt"
	"math/rand"
	"time"
)

func pull(image string) string {
	fmt.Printf("Pulling image %s\n", image)
	sleep := rand.Intn(800)
	time.Sleep(time.Millisecond * time.Duration(sleep))
	fmt.Printf("Pull done: %s\n", image)
	return image
}
