package main

import "fmt"

type IPhone struct {
	Type           string
	Memory         int
	Battery        int
	SimCardSlot    int
	MemoryCardSlot bool
}

func NewIPhone() *IPhone {
	return &IPhone{
		Type:        IPhoneType,
		Memory:      512,
		Battery:     3265,
		SimCardSlot: 1,
	}
}

func (p IPhone) GetType() string {
	return p.Type
}

func (p IPhone) Info() {
	fmt.Printf("Type: %s Memory: %d Battery: %d SimCardSlot: %d MemoryCardSlot: %v\n", p.Type, p.Memory, p.Battery, p.SimCardSlot, p.MemoryCardSlot)
}
