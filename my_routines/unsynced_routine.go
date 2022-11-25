package myroutines

import (
	"fmt"
	"sync"
)

/*
this is an unsynced routine that running who is faster
*/
var wg = sync.WaitGroup{}

func AntRoutins() { // will use ths func as main in this package
	fmt.Println("my routine package")

	wg.Add(1) // add what to wait there
	go sayHi()

	wg.Add(1)
	doRace()
	wg.Wait()

	IamSynced()
}

func sayHi() {
	fmt.Println("go go now")
	wg.Done() // remove the 1 from weight
}

// to see all race conditions and mem liaks ->  go run -race main.go  
func doRace() {
	// see the race contrition
	a := "text1"
	go func() {
		fmt.Println(a) // other
	}()
	a = "other"
}
