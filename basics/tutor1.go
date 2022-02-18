package main

import (
	"fmt"
	"reflect"
)

/*
	go run filename.go <- run in place
	go build file.go <- create .exe file
*/

func main() {
	base1()

	newData()
	newData2()
}

func base1() {
	fmt.Println("from base1 ______________________________")

	// text := "hello data"
	var text string = "some text"
	// var num int // empty value -> 0
	// var b = true

	/*
		int8 , int16 ... <- unt sizing
		uint > 0
		float16 , float32 ...
		bool

		rune +- int  size

	*/

	// fmt.Println(text)
	fmt.Println(reflect.TypeOf(text))
	// fmt.Println(num)
	// fmt.Println(b)

	x, y, z := 1, 2, 3
	fmt.Println(x, y, z)
}

var new1, new2 int = 1, 1

func newData() {
	fmt.Println("new data ____________")
	new1 := 55  // init locally 
	new2 = 77 // assign other data to global 
	fmt.Println(new1, new2)
}

func newData2() {
	fmt.Println(new1, new2)
}

