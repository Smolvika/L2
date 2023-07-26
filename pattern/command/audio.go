package main

type Audio struct {
	command Command
}

func NewAudio(command Command) *Audio {
	return &Audio{
		command: command,
	}
}

func (a *Audio) Press() {
	a.command.Execute()
}
