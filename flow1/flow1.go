package main

import (
	"fmt"
)

func main() {
	newText := loadMessage("some message", 18)
	fmt.Println(newText)

	ifMessage, ifBool := ifCheck(35)
	fmt.Println(fmt.Sprintf("text: %s , value %v", ifMessage, ifBool))

}

func loadMessage(text string, num int) string {
	// use C style text concat
	return fmt.Sprintf("hello there %s, with num %d", text, num)
}

func ifCheck(num int) (string, bool) {
	if num < 30 {
		return "text 1", false
	} else if num >= 30 && num < 45 {
		return "text2", true
	} else {
		return "text3", true
	}
}
