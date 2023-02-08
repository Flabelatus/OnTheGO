package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// reference types -> channels
// Channels are means to allow your program to send information
// from one part to another part. They are used for Go routines
// They allow for concurrent programming

// declare the channel variable
var KeyPressChan chan rune

func main() {
	KeyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key or q to quit")

	// open the keyboard
	_ = keyboard.Open()

	// closing the keyboard
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		// send the information of char to the channel
		KeyPressChan <- char

	}
}

func listenForKeyPress() {
	for {
		// get the infromation from the channel
		key := <-KeyPressChan
		fmt.Println("You pressed", string(key))
	}
}

// func main() {
// 	// call the function as a go routing with go keyword
// 	go doSomething("Hello World")

// 	fmt.Println("Do something else here")
// 	for {
// 		// do nothing
// 		// here the routine is running in the background in an infinite loop
// 	}
// }

// func doSomething(s string) {
// 	count := 0
// 	for {
// 		fmt.Println("s is", s)
// 		count = count + 1
// 		if count >= 10 {
// 			break
// 		}
// 	}
// }
