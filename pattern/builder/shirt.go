package main

import "fmt"

type Shirt struct {
	Brand    string
	Color    string
	Material string
	Length   int
	Size     int
}

func (s Shirt) InfoShirt() {
	fmt.Printf("Brand: %s Color: %s Material: %s Length: %d Size: %d\n", s.Brand, s.Color, s.Material, s.Length, s.Size)
}
