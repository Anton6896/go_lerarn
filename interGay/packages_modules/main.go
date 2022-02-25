package main

import (
	shape_pack "ant_test/shape"
	"fmt"
	myTime "ant_test/utils"
)

func main() {
	check1()
}

func check1() {
	c1 := shape_pack.NewCircle(2)
	c1.ShowRadius()
	c1.SetRadius(5)
	c1.ShowRadius()
	c1Aria := c1.Area()
	fmt.Println(c1Aria)

	myTime.GetTime()
}
