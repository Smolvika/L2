package main

import "fmt"

const (
	IPhoneType  = "iPhone"
	AndroidType = "Android"
)

type Telephone interface {
	GetType() string
	Info()
}

func New(typeTelephone string) Telephone {
	switch typeTelephone {
	case IPhoneType:
		return NewIPhone()
	case AndroidType:
		return NewAndroid()
	default:
		fmt.Printf("%s this is a non-existent type\n", typeTelephone)
		return nil
	}
}
