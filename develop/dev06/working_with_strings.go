package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run(f *Flags) {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		words := strings.Split(str, f.delimiter)

		if len(words) == 1 && f.separated {
			continue
		}
		res := make([]string, 0)
		for i := 0; i < len(f.setColumns); i++ {
			if len(words) >= f.setColumns[i] && f.setColumns[i]-1 >= 0 {
				res = append(res, words[f.setColumns[i]-1])
			} else {
				fmt.Println("there is no such field")
				os.Exit(1)
			}
		}
		fmt.Println(strings.Join(res, " "))
	}

}
