package myroutines

import (
	"fmt"
	"sync"
)

/* 
this is an unsynced routine that running who is faster 
*/
var wg = sync.WaitGroup{}

func AntRoutins() {
	fmt.Println("my routine package")

	wg.Add(1) // add what to wait there
	go sayHi()

	// wg.Add(1)
	// doRace()
	wg.Wait()

	roo()
}

func sayHi() {
	fmt.Println("go go now")
	wg.Done() // remove the 1 from weight
}

func doRace() {
	// see the race contrition
	a := "text1"
	go func() {
		fmt.Println(a) // other
	}()
	a = "other"
}
