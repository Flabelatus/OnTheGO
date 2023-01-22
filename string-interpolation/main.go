package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

// make the variable package level variable
var reader *bufio.Reader

// define custom type similar to Class
type User struct {
	UserName       string
	Age            int
	FavoriteNumber float64
	OwnsADog       bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User
	user.UserName = readStr("What is yout name?")
	user.Age = readInt("How old are you?")
	user.FavoriteNumber = readNumber("What is your favorite number?")
	user.OwnsADog = readBoolean("Do you own a dog? (y/n)")
	// Using this method and placeholders is called string interpolation
	fmt.Printf("Your name is %s. You are %d years old. Your favorite number is %.2f. Own a dog? %t\n",
		user.UserName,
		user.Age,
		user.FavoriteNumber,
		user.OwnsADog)
}

func prompt() {
	fmt.Print("> ")
}

func readStr(s string) string {
	for {
		fmt.Println(s)
		prompt()

		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please enter a value")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		// convert String to Int
		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else {
			if num < 3 {
				fmt.Println("Are you a Baby?")
			} else {
				if num > 100 {
					fmt.Println("No WAY!")
				} else {
					return num
				}
			}
		}
	}
}

func readNumber(s string) float64 {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		// convert String to Int
		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			return num
		}
	}
}

func readBoolean(s string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Println("Please type y or n")
		} else if string(char) == "y" || string(char) == "Y" {
			return false
		} else if string(char) == "n" || string(char) == "N" {
			return true
		}
	}
}
