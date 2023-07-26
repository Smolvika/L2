package main

import "fmt"

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (i *NoItemState) AddItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *NoItemState) RequestItem() error {
	return fmt.Errorf("Item out of stock ")
}

func (i *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock ")
}

func (i *NoItemState) DispenseItem() error {
	return fmt.Errorf("Item out of stock ")
}
