package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	name := "Hello World"
	fmt.Println(name)
	fmt.Println()

	fmt.Println("Bytes")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x ", name[i])
	}

	fmt.Println()
	fmt.Println()

	fmt.Println("Index\tRune\tString")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}
	fmt.Println()
	fmt.Println("Three ways to concatenate strings")

	h := "Hello, "
	w := "World"

	// using Plus sign
	myString := h + w
	fmt.Println(myString)

	//using fmt (far more efficient that using + sign)
	myString = fmt.Sprintf("%s%s", h, w)
	fmt.Println(myString)

	// using stringbuilder (the most efficient way to join strings)
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println(sb.String())

	fmt.Println()
	// slicing strings (Somehow similar to Python)
	name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println("Getting a substring")
	fmt.Println(name[10:13])

}
