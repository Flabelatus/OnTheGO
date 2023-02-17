package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "This is a clear example why we search in one case only."

	searchString := strings.ToLower(myString)

	if strings.Contains(searchString, "this") {
		fmt.Println("Found it.")

	} else {
		fmt.Println("Did not find it.")
	}

	fmt.Println(strings.ToLower(myString))
	fmt.Println(strings.ToUpper(myString))
	fmt.Println(strings.Title(myString))

	searchString = strings.Title(searchString)
	if strings.Contains(searchString, "Example") {
		searchString = strings.Replace(searchString, "Example", "EXAMPLE", 1)
	} else {
		fmt.Println("Example not found")
	}

	fmt.Println(searchString)
}
