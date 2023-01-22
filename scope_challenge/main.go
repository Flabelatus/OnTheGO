package main

import (
	"myapp/packageone"
)

var myVar = "my var,"

func main() {
	var blockVar = "main function var, "
	packageone.PrintMe(myVar, blockVar, packageone.PackageVar)
}
