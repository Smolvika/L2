package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float32
}

func (c Circle) GetPerimeter() {
	per := 2 * math.Pi * c.r
	fmt.Printf("Perimeter circle: %v\n", per)
}

func (c Circle) Accept(v Visitor) {
	v.visitForCircle(c)
}
