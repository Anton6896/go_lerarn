package main

import (
	"errors"
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
	// wright to file logs
	// file, err := os.OpenFile("flow1/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	//     log.Fatal(err)
	// }

	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	// newText := loadMessage("some message", 18)
	// fmt.Println(newText)

	ifMessage, ifBool, ifErr := ifCheck(75)
	if ifErr == nil {
		fmt.Println(fmt.Sprintf("text: %s , value %v", ifMessage, ifBool))
	} else {
		log.Fatal(ifErr)
	}

}

// func loadMessage(text string, num int) string {
// 	// use C style text concat
// 	return fmt.Sprintf("hello there %s, with num %d", text, num)
// }

func ifCheck(num int) (string, bool, error) {
	if num < 30 {
		InfoLogger.Println("got lower than 30")
		return "text 1", false, nil
	} else if num >= 30 && num < 45 {
		InfoLogger.Println("got higher than 30")
		return "text2", true, nil
	} else {
		ErrorLogger.Println("got some error")
		return "text", false, errors.New("some new error")
	}
}
