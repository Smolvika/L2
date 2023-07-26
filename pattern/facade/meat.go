package main

import "fmt"

type meat struct {
	name string
}

func newMeat(name string) *meat {
	return &meat{
		name: name,
	}
}

func (m *meat) put() {
	if m.name != "" {
		fmt.Printf("We put slices of %s on bread\n", m.name)
	}
}
