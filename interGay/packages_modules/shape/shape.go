package shape_pack

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 { // implement interface Shape
	return float32(math.Pow(float64(c.radius), 2)) * math.Pi
}

func (c *Circle) SetRadius(r float32) {
	c.radius = r
}

func (c Circle) ShowRadius() {
	fmt.Println("radius :", c.radius)
}

func NewCircle(r float32) *Circle {
	return &Circle{r}
}
