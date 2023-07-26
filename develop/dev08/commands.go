package main

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"strconv"
	"strings"
)

func executeCommands(commands []string) {
	for _, command := range commands {
		c := strings.Fields(command)
		switch c[0] {
		case "cd":
			cdCommands(c)
		case "pwd":
			pwdCommands()
		case "echo":
			echoCommands(c)
		case "kill":
			killCommands(c)
		case "ps":
			psCommands()
		case "quit":
			fmt.Fprintln(os.Stdout, "quit")
			os.Exit(0)
		default:
			fmt.Println("There is no such command")

		}
	}

}

func cdCommands(command []string) {
	if len(command) != 2 {
		fmt.Fprintln(os.Stdout, "specify the directory")
		return
	}

	if err := os.Chdir(command[1]); err != nil {
		fmt.Fprintln(os.Stdout, err)
	}
}

func pwdCommands() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		return
	}
	fmt.Println(dir)

}

func echoCommands(command []string) {
	for i := 1; i < len(command); i++ {
		fmt.Print(command[i], " ")
	}
	fmt.Println()
}

func killCommands(command []string) {
	pId, err := strconv.Atoi(command[1])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		return
	}
	process, err := os.FindProcess(pId)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		return
	}

	if err = process.Kill(); err != nil {
		fmt.Fprintln(os.Stdout, err)
	}

}

func psCommands() {
	processes, _ := ps.Processes()

	for _, process := range processes {
		fmt.Printf("process name: %v process id: %v id of the parent process: %v\n", process.Executable(), process.Pid(), process.PPid())
	}
}
