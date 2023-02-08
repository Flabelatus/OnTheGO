package main

import "fmt"

// Aggregate types (array, struct)
// You can declare an array by the var arrayName [length]type

func main() {
	var myStrings [3]string

	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println("First element in the array is", myStrings[0])

}
