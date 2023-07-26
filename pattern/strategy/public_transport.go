package main

import "fmt"

type PublicTransportStrategy struct{}

func (r *PublicTransportStrategy) Route(start, end int) {
	avgSpeed := 40
	total := end - start
	totalTime := total / avgSpeed
	fmt.Printf("PublicTransport Start:%d, End:%d Avg speed:%d Total:%d Total time: %d min \n", start, end, avgSpeed, total, totalTime)

}
