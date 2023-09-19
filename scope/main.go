package main

import "myapp/packageone"

var myvar = "I am myVar"

func main() {
	var blockVar = "I am blockVar"

	packageone.PrintMe(myvar, blockVar)
}
