package main

import (
	"fmt"
	"math"
)

// you can also combine couple interfaces

type Shape interface {
	Area() float32 // function
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	/* with that Circle will sign assignment to interface Shape  */
	return float32(math.Pow(float64(c.radius), 2)) * math.Pi
}

func main() {
	circle1 := Circle{3}
	printArea(circle1)

	// using empty interface
	printInterface("string") // works ok
	printInterface(234)      // works ok

	// and other types, look at Printf function arguments it also uses same empty interface
	interfaceCarousel(345)
}

// using interface
func printArea(s Shape) {
	fmt.Println(s.Area())
}

// using empty interface allows to push any thing in it
// ! all object is implements empty interface

func printInterface(i interface{}) {
	fmt.Printf("%+v\n", i)
}

func interfaceCarousel(i interface{}) {
	switch value := i.(type) {
	case int:
		fmt.Println("int", value)
	case string:
		fmt.Println("string", value)
	default:
		fmt.Println("default")
	}
	
	// casting interface to some type and check if cast ok 
	val, ok := i.(string) // in this case its int casting to string
	if ok {
		fmt.Println("cast ok ", val)
	}else {
		fmt.Println("cast is not ok")
		// must be user  := strconv.Itoa(12) 
	}

}
