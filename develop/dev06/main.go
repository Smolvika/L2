package main

import (
	"fmt"
	"os"
)

func main() {
	flags, err := ParseFlags()
	if err != nil {
		fmt.Println("Error parsing fields -f:", err)
		os.Exit(1)
	}
	Run(flags)
}
