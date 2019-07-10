package docker

import (
	"fmt"
	"time"
)

type networkFunction func(string) string

func measureTime(nf networkFunction) networkFunction {
	return func(argument string) string {
		start := time.Now()
		result := nf(argument)
		fmt.Printf("It took: %v\n", time.Since(start))
		return result
	}
}
