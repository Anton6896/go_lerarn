package myTime

import (
	"fmt"
	"time"
)

func GetTime() {
	currentTime := time.Now()

	fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))
	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5"))
}

/* 
	file name myUtils is not important 
	myTime "ant_test/utils" <- in main call to this func 
*/