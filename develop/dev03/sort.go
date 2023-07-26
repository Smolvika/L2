package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Sort(data [][]string, f *Flags) {
	switch {
	case f.n && !f.r:
		sort.Slice(data, func(i, j int) bool {
			int1, _ := strconv.Atoi(data[i][f.k-1])
			int2, _ := strconv.Atoi(data[j][f.k-1])
			return int1 < int2
		})
	case f.r && !f.n:
		sort.Slice(data, func(i, j int) bool {
			return data[i][f.k-1] > data[j][f.k-1]
		})
	case f.r && f.n:
		sort.Slice(data, func(i, j int) bool {
			int1, _ := strconv.Atoi(data[i][f.k-1])
			int2, _ := strconv.Atoi(data[j][f.k-1])
			return int1 > int2
		})
	default:
		sort.Slice(data, func(i, j int) bool {
			return data[i][f.k-1] < data[j][f.k-1]
		})
	}
}

func RemoveDuplicate(data [][]string) [][]string {
	duplicate := make(map[string]int)
	for _, val := range data {
		str := strings.Join(val, " ")
		_, ok := duplicate[str]
		if !ok {
			duplicate[str]++
		} else {
			duplicate[str]--
		}
	}
	res := make([][]string, 0)
	for key, val := range duplicate {

		if val == 1 {
			res = append(res, strings.Fields(key))
		}
	}

	return res
}

func ChangedResult(name string, data [][]string) error {
	f, err := os.OpenFile(name, os.O_WRONLY, 0777)
	defer f.Close()

	if err != nil {
		return err
	}

	for _, val := range data {
		if _, err = f.WriteString(fmt.Sprintf("%s\n", val)); err != nil {
			return err
		}
	}

	return nil
}
