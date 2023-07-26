package main

func main() {
	start := 1
	end := 110
	strategies := []Strategy{&PublicTransportStrategy{}, &RoadStrategy{}, &WalkStrategy{}}
	nav := Navigator{}
	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		nav.Route(start, end)
	}
}
