package main

import (
	"bufio"
	"errors"
	"os"
)

func OpenFile(name string) (*os.File, error) {

	if name != "" {
		file, err := os.Open(name)
		defer file.Close()

		if err != nil {
			return nil, err
		}

		return file, nil
	}
	return nil, errors.New("File name not entered")
}

func ReadFile(file *os.File) []string {
	result := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}
