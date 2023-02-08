package main

import "fmt"

// reference types -> maps

func main() {
	// maps are always created by the make keyword
	var intMap = make(map[string]int)

	// they are reference type, and it is always passed by refernece.
	// you would rarely see a pointer to a map

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	// delete a key from the map
	delete(intMap, "two")

	// test if a key is in the map
	// keyVariable, boolean := intMap[key]
	element, ok := intMap["one"]
	if ok {
		fmt.Println(element, "is in the map")
	} else {
		fmt.Println("The key you searched is not in the map")
	}

	// overwriting values of a key in the map
	intMap["one"] = 10

	for key, value := range intMap {
		fmt.Println(key, value)
	}
}
