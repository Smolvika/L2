package main

import "fmt"

type sandwich struct {
	cheese     *cheese
	vegetables *vegetables
	meat       *meat
}

func newSandwich(cheese string, meat string, vegetables string) *sandwich {
	return &sandwich{
		cheese:     newCheese(cheese),
		meat:       newMeat(meat),
		vegetables: newVegetables(vegetables),
	}
}

func (s *sandwich) prepare() {
	fmt.Println("Starts making a sandwich")
	s.cheese.put()
	s.vegetables.put()
	s.meat.put()
	fmt.Println("The sandwich is ready")
}
