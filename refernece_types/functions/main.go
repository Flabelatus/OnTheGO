package main

import "fmt"

// reference types -> functions

type Animal struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

// a receiver function (that links to a type similar to methods in classes)
func (a *Animal) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
	fmt.Println()
}

func main() {
	// myTotal := sumMany(2, 20, 23, 1, 23, 120, -200)
	// fmt.Println(myTotal)

	var dog Animal
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.NumberOfLegs = 4
	dog.Says()

	cat := Animal{
		Name:         "Bob the cat",
		Sound:        "Meouw",
		NumberOfLegs: 4,
	}
	cat.Says()
}

// // declare a simple function
// func addNumber(x, y int) int {
// 	return x + y
// }

// a function with variatic parameter
// If a you have a function with multiple parameters and one of them
// can be a variatic, that needs to be placed at the end of the parameters
func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}

	return total
}
