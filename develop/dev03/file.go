package main

import (
	"bufio"
	"os"
	"strings"
)

func OpenAndReadFile(name string) ([][]string, error) {
	dataFile := make([]string, 0)

	if name != "" {
		file, err := os.Open(name)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			dataFile = append(dataFile, scan.Text())
		}
	}
	res := make([][]string, 0)

	for _, val := range dataFile {
		res = append(res, strings.Fields(val))
	}

	return res, nil
}
