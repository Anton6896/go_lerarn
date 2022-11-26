package myroutines

import (
	"fmt"
	"sync"
)

var counter = 0
var m = sync.RWMutex{}

// var wg = sync.WaitGroup{} //get it from mr.go

func IamSynced() {
	// runtime.GOMAXPROCS(20) // unsafe working with golang ! can bring memory overhanded 
	for i := 0; i < 10; i++ {
		wg.Add(2)

		m.RLock() // read lock
		go readIt()

		m. Lock() // wright lock 
		go writeTo()
	}
	wg.Wait() // wight for all 
}

func readIt() {
	fmt.Printf("number : %v\n", counter)
	m.RUnlock()
	wg.Done()

}

func writeTo() {
	counter++
	m.Unlock()
	wg.Done()
}
