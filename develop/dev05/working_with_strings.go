package main

import (
	"fmt"
	"regexp"
	"strings"
)

type SearchData struct {
	str     []string
	pattern string
	flags   *Flags
	re      *regexp.Regexp
}

type SearchResources struct {
	str     []string
	c       *Flags
	re      *regexp.Regexp
	pattern string
}

func NewSearchData(str []string, pattern string, flags *Flags, re *regexp.Regexp) *SearchData {
	return &SearchData{
		str:     str,
		pattern: pattern,
		flags:   flags,
		re:      re,
	}
}

func (s SearchData) SearchString() {
	var count int
	for i := 0; i < len(s.str); i++ {
		coincidence, notExclude := s.ValidateString(i)
		if coincidence == true && s.flags.count {
			count++
		}

		if notExclude {
			if s.flags.num {
				fmt.Printf("Строка: %d\n", i+1)
			}
			s.BeforeString(i, "before")
			s.AfterString(i, "after")
			s.ContextString(i)
		}

	}
	if s.flags.count {
		fmt.Printf("Количество строк: %d\n", count)
	}

}

func (s SearchData) ValidateString(i int) (bool, bool) {
	var ans bool
	if s.flags.fixed {
		if s.flags.ignore {
			ans = strings.ToLower(s.str[i]) == strings.ToLower(s.pattern)
		} else {
			ans = s.str[i] == s.pattern
		}
	} else {
		ans = s.re.MatchString(s.str[i])
	}

	return ans, ans != s.flags.invert
}

func (s SearchData) BeforeString(i int, str string) {
	var param int
	if str == "before" {
		param = s.flags.before
	} else {
		param = s.flags.context
	}

	if param == 0 {
		return
	}

	var start int
	if i-param > 0 {
		start = i - param
	}

	for j := start; j < i; j++ {
		fmt.Println(s.str[j])
	}
	fmt.Println()
}

func (s SearchData) AfterString(i int, str string) {
	var param int
	if str == "after" {
		param = s.flags.after
	} else {
		param = s.flags.context
	}

	if param == 0 {
		return
	}
	end := len(s.str) - 1
	if i+param < len(s.str)-1 {
		end = i + param
	}

	for j := i + 1; j <= end; j++ {
		fmt.Println(s.str[j])
	}
	fmt.Println()
}

func (s SearchData) ContextString(i int) {
	if s.flags.context != 0 {
		s.BeforeString(i, "context")
		s.AfterString(i, "context")
	}
}
