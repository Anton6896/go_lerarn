package main

import (
	"fmt"
	"strconv"
	"reflect"
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
		car: car{mark: "honda", speed: 100},
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

// continue at 2.17
func main() {
	myStructs()
}
