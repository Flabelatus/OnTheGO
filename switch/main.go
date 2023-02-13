package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	userInput, _ := reader.ReadString('\n')

	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)
	
	// if/else statement

	// if userInput == "Hello" {
	// 	fmt.Println("Hello to you as well")
	// } else if userInput == "Bye" {
	// 	fmt.Println("Good Bye")
	// } else if userInput == "Hello World" {
	// 	fmt.Println("Good Bye World")
	// } else {
	// 	fmt.Println(userInput)
	// }

	// Switch syntax instead of if/else syntax
	switch userInput {
	case "Hello":
		fmt.Println("Hello to you as well")
	case "Bye":
		fmt.Println("Good Bye")
	case "Hello World":
		fmt.Println("Good Bye world")
	default:
		fmt.Print(userInput)
	}
}
