package main

import "fmt"

type WalkStrategy struct{}

func (r *WalkStrategy) Route(start, end int) {
	avgSpeed := 4
	total := end - start
	totalTime := total / avgSpeed
	fmt.Printf("Walk Start:%d, End:%d Avg speed:%d Total:%d Total time: %d min \n", start, end, avgSpeed, total, totalTime)

}
