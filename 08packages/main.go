package main

import (
	"fmt"
	"math"
	"github.com/Brady-Higgins/example/shapes"
)

func main() {
	fmt.Println(math.Pi)

	circle := shapes.Circle{Radius: 2.0}
	fmt.Println(circle.Radius)
}