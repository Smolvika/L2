package main

import "fmt"

type RoadStrategy struct{}

func (s *RoadStrategy) Route(start, end int) {
	avgSpeed := 60
	trafficJam := 2
	total := end - start
	totalTime := trafficJam * total / avgSpeed
	fmt.Printf("Road Start:%d, End:%d Avg speed:%d Traffic jam:%d Total:%d Total time: %d min \n", start, end, avgSpeed, trafficJam, total, totalTime)
}
