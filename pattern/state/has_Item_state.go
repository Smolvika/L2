package main

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (h *HasItemState) AddItem(count int) error {
	h.vendingMachine.incrementItemCount(count)
	fmt.Printf("%d items added\n", count)
	return nil
}

func (h *HasItemState) RequestItem() error {
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
		return fmt.Errorf("No item present ")
	}
	fmt.Printf("Item requested\n")
	h.vendingMachine.setState(h.vendingMachine.itemRequested)
	return nil
}

func (h *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("Please select item first ")
}

func (h *HasItemState) DispenseItem() error {
	return fmt.Errorf("Please select item first ")
}
