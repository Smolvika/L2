package main

import "fmt"

type Android struct {
	Type           string
	Memory         int
	Battery        int
	SimCardSlot    int
	MemoryCardSlot bool
}

func NewAndroid() *IPhone {
	return &IPhone{
		Type:           AndroidType,
		Memory:         64,
		Battery:        4200,
		SimCardSlot:    2,
		MemoryCardSlot: true,
	}
}

func (p Android) GetType() string {
	return p.Type
}

func (p Android) Info() {
	fmt.Printf("Type: %s Memory: %d Battery: %d SimCardSlot: %d MemoryCardSlot: %v\n", p.Type, p.Memory, p.Battery, p.SimCardSlot, p.MemoryCardSlot)
}
