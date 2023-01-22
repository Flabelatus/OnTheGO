package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		mainLine()
		var reader = bufio.NewReader(os.Stdin)
		r, _ := reader.ReadString('\n')
		r = strings.Replace(r, "\r\n", "", -1)

		if r == "q" {
			fmt.Println("Sorry to see you going!")
			break
		} else {
			fmt.Println("Hello", r, "!")
		}
	}
}

func mainLine() {
	fmt.Println("What is your name?")
}
