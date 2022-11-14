package main

import (
	"fmt"
)

// declare var as block at top level 
var (
	fname string = "block declaration"
	fage int = 45
)

func main() {
	fmt.Printf("%v %v \n", fname, fage)

	var age int
	name := "my type is"
	age = 15

	// https://pkg.go.dev/fmt
	fmt.Printf("Hello World! %v %v \n", name, age)
	fmt.Printf("%v : %T", name, name)


}

// go mod init
