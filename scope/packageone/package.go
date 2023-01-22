package packageone

// If the variable starts with small letter, it wont be available outside this 
// package.
var privateVar = "I am private"

// if the variable starts with the capital letter, It is available outside
// this package as well.
var PublicVar = "I am public (or exported)"

func notExported() {

}

func Exported() {
	
}