package main

type Shape interface {
	GetPerimeter()
	Accept(v Visitor)
}
