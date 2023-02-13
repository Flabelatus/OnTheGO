package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// seed random generator
	rand.Seed(time.Now().UnixNano())
	i := 1000

	// execute a loop while some condition is true
	for i > 100 {
		// get a random number between 1 and 1001
		i = rand.Intn(1000) + 1
		if i > 100 {
			fmt.Println("i is", i)
		}

	}

	fmt.Println("Got", i, "and broke out of loop")
}
