package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// continue at 4.22
func main() {
	functionsCheck()
}

func functionsCheck() {
	// arr1 := []int{1, 2, 3, 4}
	val := checkAll("foo", 1, 2, 3, 4)
	fmt.Println("dereference val :", *val)

	d, err := devider(1, 0)
	if err != nil {
		println("have some error")
	}
	fmt.Println("result ", d)

	f := func () {
		fmt.Println("i am inside function")
	}

	f()
}

func devider(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("zero devider")
	}

	return a / b, nil
}

func checkAll(foo string, values ...int) *int {
	// same as *args
	fmt.Printf("str: %v, args : %v\n", foo, values)
	/*
		in goleng we can return this location in memory it will be saved after function is closed too !
	*/
	res := 0
	return &res
}

func pointers() {
	var a int = 45
	var b *int = &a
	println(b)
	println(*b)
	a = 1
	println(*b) // point to same location as var a
	*b = 4
	println(a)

	type myFoo struct {
		foo string
	}
	var mf *myFoo = new(myFoo)
	(*mf).foo = "bar" // first dereference mf , then assign value to
	println(mf.foo)   // syntactic sugar , golang will diref by it self

}

func controlFlow() {
	// try catch ...

	// defer working on LIFO array
	defer println("1")
	defer println("2") // will trigger when controlFlow ends
	println("3")

	// Panic is an state the  golang cant handle this issue
	a, b := 1, 0
	ans := a / b
	println(ans)

}

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

func myMaps() {
	myMap := make(map[string]int)
	myMap = map[string]int{
		"val1": 1,
		"val2": 2,
	}
	myMap["val3"] = 3
	delete(myMap, "val2")
	fmt.Println(myMap)

	// check if in the map
	inside, ok := myMap["val2"]
	fmt.Printf("is it inside %v, %v", inside, ok)
}

func myStructs() {
	type patient struct {
		name string
		age  int
	}

	type Doctor struct {
		name       string
		profession string
		patients   []patient
	}

	docA := Doctor{
		name:       "docA",
		profession: "professionA",
		patients: []patient{
			{
				name: "patient1",
				age:  12,
			},
			{
				name: "patient2",
				age:  14,
			},
		},
	}

	fmt.Println(docA.patients[0].name)
	fmt.Println(docA.patients[1].name)

	newCar := struct {
		make    string
		model   string
		mileage int
	}{
		make:    "Ford",
		model:   "Taurus",
		mileage: 200000,
	}
	fmt.Printf("newCar %v", newCar)

	// composition relationship
	fmt.Println("composition ----------- ")
	type car struct {
		mark  string
		speed int
	}

	type honda struct {
		car
		name string
	}

	car1 := honda{
		car:  car{mark: "honda", speed: 100},
		name: "civic",
	}

	fmt.Println(car1)

	// validation | reflection
	type human struct {
		name string `required max:"100"`
	}

	h := reflect.TypeOf(human{})
	field, _ := h.FieldByName("name")
	println(field.Tag)
}

func getTrue() bool {
	return true
}

func ifAndSwitch() {
	num1 := 4
	if num1 > 10 || getTrue() {
		println("inside the if statement")
	} else if getTrue() {
		println("the else in work")
	}

	switch num1 {
	case 1, 3:
		println("some num")
	case 4:
		println("this is 4")
	default:
		println("this is default")
	}

	switch {
	case num1 > 3:
		println("its big")
	default:
		println("i am default")
	}

	var i interface{} = "1" // interface can be any type of (int, str, ...)
	switch i.(type) {
	case int:
		println("int")
		break // its working ok
		println("break still working")
	case bool:
		println("bool")
	case [2]int:
		println("array with 2 ints")
	default:
		println("some other")
	}
}

func looping() {
	for i := 0; i < 4; i++ {
		print("num: %v, ", i)
		if i == 3 {
			println()
		}
	}
	println("other line")

	i := 0
	for i < 3 {
		i++
		println("some job")
	}

	for {
		i++
		if i > 5 {
			break
		}
		if i == 4 {
			continue
		}
		println("some while loop")
	}

	// loop labeling (great new feature for me)

Loop:
	for i := 0; i < 4; i++ {
		println("mainloop")
		for j := 0; j < 4; j++ {
			println("inner loop")
			if j == 2 {
				break Loop
			}
		}
	}

	myArr := []int{1, 2, 3, 4}
	for k, v := range myArr {
		println(k, v)
	}

	myMap := map[string]int{
		"val1": 1,
		"val2": 2,
	} // have no order !

	for _, v := range myMap {
		println(v)
	}

}
