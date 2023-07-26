package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	for {
		if scan.Scan() {
			command := scan.Text()
			commands := strings.Split(command, "|")
			executeCommands(commands)
		}
	}
}
