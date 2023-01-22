package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

// func main() {

// 	// var thingToSay string
// 	// thingToSay = "Hello World"

// 	thingToSay := "Hello World" // Short cut for assigning variables
// 	sayHelloWorld(thingToSay)
// }

// func sayHelloWorld(thePhrase string) {
// 	fmt.Println(thePhrase)
// }

func main() {
	// reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	for {
		fmt.Print(">")
		userInput, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		// Strip off the entered key from the string
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}
	}
}
