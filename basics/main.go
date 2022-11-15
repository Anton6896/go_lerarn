package main

import (
	"fmt"
	"strconv"
)

func basics() {
	// declare var as block at top level
	var (
		fname string = "block declaration"
		fage  int    = 45
	)
	const (
		constA = iota // will use this pattern for other const vars
		constB
		constC
	)

	fmt.Printf("%v %v \n", fname, fage)

	var age int
	name := "my type is"
	age = 15

	// https://pkg.go.dev/fmt
	fmt.Printf("Hello World! %v %v \n", name, age)
	fmt.Printf("%v : %T\n", name, name)

	n := 15
	fmt.Printf("int val : %v is %T\n", n, n)

	c := strconv.Itoa(n)
	fmt.Printf("converted val : %v is %T\n", c, c)

	// go working with string as an collection of bytes
	text := "this is the string"
	bText := []byte(text)
	fmt.Printf("this is the string : %v : %T\n", bText, bText) //utf8

	fmt.Printf("aiota : %v ", constA)
	fmt.Printf("aiota : %v ", constB)
	fmt.Printf("aiota : %v \n", constC)
}

func great1() {
	const (
		positionA = iota // can be used as example as combined keys
		positionB
		positionC

		sideA
		sideB
		sideC
	)

	roleA := positionA | sideA | sideB
	// roleB :=positionB | sideB | sideC

	fmt.Printf("inpositionA %v \n", positionA&roleA == positionA)
}

func main() {
	great1()

}
