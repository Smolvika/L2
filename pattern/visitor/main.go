package main

import "fmt"

func main() {
	square := Square{a: 4.9}
	circle := Circle{r: 3.1}
	calculator := &areaCalculator{}

	square.GetPerimeter()
	square.Accept(calculator)
	fmt.Printf("Area square: %v\n", calculator.area)

	circle.GetPerimeter()
	circle.Accept(calculator)
	fmt.Printf("Area circle: %v\n", calculator.area)
}
