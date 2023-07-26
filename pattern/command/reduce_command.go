package main

type ReduceCommand struct {
	device Device
}

func NewReduceCommand(device Device) *ReduceCommand {
	return &ReduceCommand{
		device: device,
	}
}

func (c *ReduceCommand) Execute() {
	c.device.Reduce()
}
