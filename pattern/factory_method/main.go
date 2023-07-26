package main

func main() {
	types := []string{IPhoneType, AndroidType, "_type"}
	for _, t := range types {
		tele := New(t)
		if tele != nil {
			tele.Info()
		}
	}
}
