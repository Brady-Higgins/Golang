package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

type Shape interface {
	area() float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (r *Rectangle) area() float64 {
	return r.Width * r.Height
}


func main() {
	c := Circle{Radius: 4.0}
	r := Rectangle{Width: 2.0, Height: 3.0}

	shapes := []Shape{&c, &r}
	fmt.Println(shapes[0].area())
}