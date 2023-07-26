package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type Unpacker interface {
	unpack() (string, error)
}

type Sequence struct {
	str string
}

func (s Sequence) unpack() (string, error) {
	var num string
	var letter rune
	var builder strings.Builder

	if s.str == "" {
		return "", nil
	}
	str := []rune(s.str)
	if unicode.IsDigit(str[0]) {
		return "", errors.New("invalid string")
	}

	letter = str[0]
	for i := 1; i < len(str); i++ {

		if unicode.IsDigit(str[i]) {
			num += string(str[i])
		} else {

			n, _ := strconv.Atoi(num)
			if n == 0 {
				n++
			}
			builder.WriteString(strings.Repeat(string(letter), n))

			num = ""
			letter = str[i]
		}
	}

	n, _ := strconv.Atoi(num)
	if n == 0 {
		n++
	}
	builder.WriteString(strings.Repeat(string(letter), n))

	return builder.String(), nil
}

type EscapeSequence struct {
	str string
}

func (s EscapeSequence) unpack() (string, error) {
	var num string
	var letter rune
	var builder strings.Builder
	str := []rune(s.str)

	for i := 0; i < len(str); i++ {
		if string(str[i]) != "\\" {
			builder.WriteRune(str[i])
		} else {
			if i+1 < len(str) {

				letter = str[i+1]

				i = i + 2
				for i < len(str) {
					if unicode.IsDigit(rune(s.str[i])) {
						num += string(s.str[i])
					} else {
						break
					}
					i++
				}

				i--

				n, _ := strconv.Atoi(num)
				if n == 0 {
					n++
				}

				builder.WriteString(strings.Repeat(string(letter), n))
				num = ""
			}

		}

	}
	return builder.String(), nil
}
