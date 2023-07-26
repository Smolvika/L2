package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var str string
	fmt.Fscan(in, &str)
	var s Unpacker
	if strings.Contains(str, "/") {
		s = EscapeSequence{str: str}
	} else {
		s = Sequence{str: str}
	}
	str, err := s.unpack()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(str)
	}
}
