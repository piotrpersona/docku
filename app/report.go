package app

import (
	"fmt"
	"time"
)

func report(start time.Time) {
	fmt.Println("Done!")
	fmt.Printf("In time: %v\n", time.Since(start))
}
