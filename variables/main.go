package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "press Enter when ready."

var firstNumber int
var secondNumber int
var subtraction int
var answer int

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	firstNumber = rand.Intn(8) + 2
	secondNumber = rand.Intn(8) + 2
	subtraction = rand.Intn(8) + 2
	answer = firstNumber*secondNumber - subtraction

	rest(firstNumber, secondNumber, subtraction, answer)
}

func rest(firstNumber, secondNumber, subtraction, answer int) {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("Start Here!")
	fmt.Println("___________")
	fmt.Println("")
	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now Multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you thought", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// Give the answer
	fmt.Println("The answer is", answer)
	time.Sleep(2 * time.Second)
}
