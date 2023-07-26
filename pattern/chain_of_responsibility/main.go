package main

import "fmt"

func main() {
	handlers := &Handler1{Next: &Handler2{Next: &Handler3{}}}
	fmt.Println(handlers.SendRequest(1))
}
