package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3} // res C array (closed in space)
	fmt.Println(arr1)

	// var slice1 []int  // actual pointer to space of array
	// fmt.Println(slice1)
	// slice1[0] = 1  // will provide error <- slise1 is nil

	slice2 := make([]int, 3) // make will create array on allocated space 
	slice2[0] = 1
	fmt.Println(slice2)
}
