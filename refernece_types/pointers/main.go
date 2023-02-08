package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

func main() {
	x := 10

	// using the & sign to get the pointer to the specific variable location
	// in the memory. This is called referencing to a pointer
	myFirstPointer := &x

	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	// when you want to print or change the value of the variable to which my pointer is pointing at
	// in the memory we use *
	*myFirstPointer = 15
	fmt.Println("x is now", x)

	changeValueOfPointer(&x)
	fmt.Println("After the function call the x is now", x)
}

// an example of how you can change the value of a varible without having it
// in the scope of the function using the pointers.
func changeValueOfPointer(num *int) {
	*num = 25
}
