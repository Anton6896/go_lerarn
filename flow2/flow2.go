package main

import (
	"fmt"
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	text := "change me"

	fmt.Println("before change: " + text)

	InfoLogger.Println("referencing data")
	textChanger(&text) // pass memory location <- referesing
	fmt.Println("after change: " + text)

	number := 5
	var p *int // create pointer 
	fmt.Println(p)
	p = &number
	fmt.Println(p)  // assign place to pointer 

	*p = 11 // dereference update number 
	fmt.Println(number)

}	

/*
	as an C leng you have an ability to work with memory location
*/

func textChanger(text *string) { // ask for memory location
	InfoLogger.Println("dereferencing data")
	*text = "i am new text" // <- de refereeing
	// assign to this mem location this data
}
