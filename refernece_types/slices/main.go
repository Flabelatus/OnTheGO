package main

import (
	"fmt"
	"sort"
)

// reference types -> slices\

func main() {
	var animals []string

	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "fish")
	animals = append(animals, "horse")

	fmt.Println(animals)

	// in this for loop, _ is the index, and x i the current
	// item in the iteration
	for _, x := range animals {
		fmt.Println(x)
	}

	for i, x := range animals {
		fmt.Println(i, x)
	}

	// similar to slicing lists in Python
	fmt.Println("First two animals are", animals[:2])

	// get the length of the slice
	fmt.Println("The length of the slice is", len(animals))

	// Use the StringsAreSorted method to check if something is sorted or not
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	// you can now sort the slice (if the slice includes strings type)
	// sorting happens inplace
	sort.Strings(animals)
	fmt.Println("Is it sorted now?", sort.StringsAreSorted(animals))

	animals = DeleteFromSlice(animals, 2)
	fmt.Println(animals)
}

func DeleteFromSlice(array []string, i int) []string {
	// remove a value from the slice at an index
	// parameters are slice of strings, index as int
	// returns a slice of strings
	// this method breaks the sort of the slice
	// if you need it to be sorted, you can again use the sort.Stirngs(a)

	// assign the value at index i to become the last item in the slice
	array[i] = array[len(array)-1]
	// then make the value to empty strings
	array[len(array)-1] = ""
	// then trucate the slice from i=0 to i=len(slice) - 1
	array = array[:len(array)-1]
	// return the slice
	return array
}
