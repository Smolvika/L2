package main

type Navigator struct {
	Strategy
}

func (n *Navigator) SetStrategy(str Strategy) {
	n.Strategy = str
}
