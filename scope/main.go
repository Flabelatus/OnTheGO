package main

import (
	"fmt"
	"myapp/packageone"
)

var one = "One"

func main() {
	var somethingElse = "The block level variable"
	fmt.Println(somethingElse)
	myFunc()

	newString := packageone.PublicVar
	fmt.Println("From packageone:", newString)

	packageone.Exported()

}

func myFunc() {
	fmt.Println(one)
}
