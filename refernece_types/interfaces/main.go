package main

import "fmt"

type Animal interface {
	// methods that can be used inside interface allow them to be used
	// for various types
	Says() string
	HowManyLegs() int
}

// type Dog
type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

// methods for Dog (function with receiver to Dog)
func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

// type Cat
type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

// method for Cat (function with receiver to Cat)
func (c *Cat) Says() string {
	return c.Sound
}

func (c *Cat) HowManyLegs() int {
	return c.NumberOfLegs
}

func main() {
	// create dog
	var dog Dog
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.NumberOfLegs = 4

	// create cat
	cat := Cat{
		Name:         "cat",
		Sound:        "meouw",
		NumberOfLegs: 4,
		HasTail:      true,
	}

	Riddle(&cat)
	Riddle(&dog)

}

func Riddle(a Animal) {
	riddle := fmt.Sprintf("What animal is it? it says %s and has %d legs", a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}
