package main

import "fmt"

type vegetables struct {
	name string
}

func newVegetables(produce string) *vegetables {
	return &vegetables{name: produce}
}

func (v *vegetables) put() {
	if v.name != "" {
		fmt.Printf("We put slices of %s on bread\n", v.name)
	}
}
