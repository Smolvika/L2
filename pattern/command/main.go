package main

func main() {
	pl := NewPlayer()
	increase := NewIncreaseCommand(pl)
	reduce := NewReduceCommand(pl)
	increaseAudio := NewAudio(increase)
	increaseAudio.Press()

	reduceAudio := NewAudio(reduce)
	reduceAudio.Press()
}
