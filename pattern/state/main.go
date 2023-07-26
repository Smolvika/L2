package main

import (
	"log"
)

func main() {
	machine := NewVendingMachine(1, 200)
	if err := machine.RequestItem(); err != nil {
		log.Fatalf(err.Error())
	}

	if err := machine.InsertMoney(10); err != nil {
		log.Fatalf(err.Error())
	}
	if err := machine.DispenseItem(); err != nil {
		log.Fatalf(err.Error())
	}

	if err := machine.AddItem(10); err != nil {
		log.Fatalf(err.Error())
	}

	if err := machine.RequestItem(); err != nil {
		log.Fatalf(err.Error())
	}

	if err := machine.InsertMoney(200); err != nil {
		log.Fatalf(err.Error())
	}

	if err := machine.DispenseItem(); err != nil {
		log.Fatalf(err.Error())
	}
}
