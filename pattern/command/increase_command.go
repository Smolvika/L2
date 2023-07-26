package main

type IncreaseCommand struct {
	device Device
}

func NewIncreaseCommand(device Device) *IncreaseCommand {
	return &IncreaseCommand{
		device: device,
	}
}

func (c *IncreaseCommand) Execute() {
	c.device.Increase()
}
