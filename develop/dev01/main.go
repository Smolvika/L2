package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func getTime() (time.Time, error) {
	response, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		return time.Time{}, err
	}

	return response, nil
}

func main() {
	t, err := getTime()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Current time: %v", t)
}
