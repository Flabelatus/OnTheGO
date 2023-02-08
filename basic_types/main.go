package main

import "log"

// basic types (numbers, strings, booleans)
var myInt int
var myUint uint
var myFloat float32
var myFloat64 float64

// aggregate types (array, struct)

// reference types (Pointers, slices, maps, functions, channels)

// interface types

func main() {
	myInt = 10
	myUint = 20

	myFloat = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := ""

	myString = "Javid"
	log.Println(myString)

	// Boolean types are only True or Falese, not 0 or 1 ...
	var myBool = true
	log.Println(myBool)
}
