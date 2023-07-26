package main

import "fmt"

type Square struct {
	a float32
}

func (s Square) GetPerimeter() {
	per := 4 * s.a
	fmt.Printf("Perimeter square: %v\n", per)
}

func (s Square) Accept(v Visitor) {
	v.visitForSquare(s)
}
