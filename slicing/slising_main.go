package main

import "fmt"

func main() {
	arr1 := [3]int{1, 2, 3} // res C array (closed in space)
	fmt.Println(arr1)

	// var slice1 []int  // actual pointer to space of array
	// fmt.Println(slice1)
	// slice1[0] = 1  // ! will provide error <- slise1 is nil

	slice2 := make([]int, 3) // make will create array on allocated space
	slice2[0] = 1
	fmt.Println(slice2)
	addToSlice(slice2)
	fmt.Println(slice2)

	slice2 = append(slice2, 234) // add data to slice (if not enough space will copy to new slice underhood)
	fmt.Println(slice2)

	/*
	* to remove element from slice must implement function RemoveIndex
	 */

	matrixSlice()
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func addToSlice(data []int) {
	/* slice will put here reference to array in it ,
	so will update data in place  */
	data[0] = 11
	fmt.Println(fmt.Sprintf("added new data %d", data[0]))
}

func matrixSlice() {
	const MATRIX_SIZE int = 10

	matrix := make([][]int, MATRIX_SIZE)
	fmt.Println(fmt.Sprintf("working with matrix , with size : %d", MATRIX_SIZE))

	for i := 0; i < MATRIX_SIZE; i++ {
		for j := 0; j < MATRIX_SIZE; j++ {
			matrix[i] = make([]int, MATRIX_SIZE)
			// matrix[i][i] = i  // identity matrix
		}
		fmt.Println(matrix[i])
	}
}
