package packageone

import "fmt"

var PackageVar = "I am PackageVar from the package"

func PrintMe(myVar, blockVar string) {
	fmt.Println(PackageVar, myVar, blockVar)
}
