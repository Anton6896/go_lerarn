package main

import "fmt"

func main() {
	defer panicHandler()

	fmt.Println("mail func --------------")


	/* 
		defer will run it at end (like destractor)
		!! it will add additional time to function run 50 ms 
	*/	
	panic_func()  

	fmt.Println("some text") // ! not seen 
}	

func panic_func() {
	fmt.Println("panic func")
	arr1 := []int{1, 2, 3}
	fmt.Println(arr1)

	fmt.Println(arr1[3])  // <- error panic 
}

func panicHandler(){
	if r := recover(); r != nil {
		fmt.Println(r)
	}	

	fmt.Println("defer handling panic in code")
}