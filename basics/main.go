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

func arrays() {
	arr1 := [...]int{1, 2, 3} // ... add same as i put in initiation
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	var arr2 [3]string
	arr2[1] = "suprime"
	fmt.Println(arr2[1])

	var identityMatrix [3][3]int = [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	for i := 0; i < len(identityMatrix); i++ {
		fmt.Println(identityMatrix[i])
	}

	// array slice
	fmt.Println(identityMatrix[0][1:])
	fmt.Println(identityMatrix[1][:2])
	fmt.Println(identityMatrix[0][:2])

	// concat 2 slices
	arr3 := []int{1, 2, 3}
	arr4 := []int{4, 5, 6}
	fmt.Println(arr3)
	arr3 = append(arr3, arr4...)
	fmt.Println(arr3)


}

// continue at 2.17
func main() {
	arrays()

}
