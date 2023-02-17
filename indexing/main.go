package main

import "fmt"

func main() {
	courseName := "Learn Go for Beginners Crash Course"

	// Prints the Rune value
	fmt.Println(courseName[0])

	// to print string value wrap it in the string function
	fmt.Println(string(courseName[0]))
	fmt.Println(string(courseName[6]))

	for i := 0; i <= 21; i++ {
		fmt.Print(string(courseName[i]))
	}

	fmt.Println()

	for i := 13; i <= 21; i++ {
		fmt.Print(string(courseName[i]))
	}
}
