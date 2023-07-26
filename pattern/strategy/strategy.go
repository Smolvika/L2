package main

type Strategy interface {
	Route(start, end int)
}
