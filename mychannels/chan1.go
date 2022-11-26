package mychannels

import (
	"fmt"
	"sync"
	"time"
)

var w = sync.WaitGroup{}

// logger
const (
	logInfo    = "INFO"
	logWarning = "WARNING"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

func logger() {
	select {
	case entry := <-logCh:
		fmt.Printf("%v = [%v] %v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		doneCh <- struct{}{}
	case <-doneCh:
		break
	}

}

func MyChannel() { // main function
	go logger()
	// w.Add(1)
	logCh <- logEntry{time: time.Now(), severity: logInfo, message: "this is logger massage"}

	// chan1()
	bufferedChan()

	w.Wait()

	defer func() {
		close(doneCh)
	}()

}

// regular channel
func chan1() {
	ch := make(chan int)
	w.Add(2) // 2 tasks

	go func(ch <-chan int) { // receiver
		i := <-ch
		fmt.Println(i)
		w.Done()
	}(ch)

	go func(ch chan<- int) { // sender
		ch <- 15
		w.Done()
	}(ch)
}

// buffered
func bufferedChan() {
	ch := make(chan int, 50) // have an 50 slots
	w.Add(2)

	// consumer
	go func(ch <-chan int) {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				// if channel is closed
				break
			}
		}
		w.Done()
	}(ch)

	// producer
	go func(ch chan<- int) {
		ch <- 12
		ch <- 55
		close(ch) // totals close this channel
		w.Done()
	}(ch)
}
