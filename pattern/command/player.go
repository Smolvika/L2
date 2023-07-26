package main

import "fmt"

type Player struct{}

func NewPlayer() *Player {
	return &Player{}
}

func (p Player) Reduce() {
	fmt.Println("Reducing the sound of the player")
}

func (p Player) Increase() {
	fmt.Println("Increase the sound of the player")
}
