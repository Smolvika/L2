package main

import "fmt"

type cheese struct {
	name string
}

func newCheese(name string) *cheese {
	return &cheese{
		name: name,
	}
}

func (c *cheese) put() {
	if c.name != "" {
		fmt.Printf("We put slices of %s cheese on bread\n", c.name)
	}
}
