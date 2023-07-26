package main

import (
	"math"
)

type Visitor interface {
	visitForSquare(Square)
	visitForCircle(Circle)
}
type areaCalculator struct {
	area float32
}

func (a *areaCalculator) visitForSquare(s Square) {
	a.area = s.a * s.a
}

func (a *areaCalculator) visitForCircle(c Circle) {
	a.area = 2 * math.Pi * c.r
}
