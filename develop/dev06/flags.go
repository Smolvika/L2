package main

import (
	"flag"
	"sort"
	"strconv"
	"strings"
)

type Flags struct {
	fields     string
	delimiter  string
	setColumns []int
	separated  bool
}

func ParseFlags() (*Flags, error) {
	var flags Flags
	flag.StringVar(&flags.fields, "f", "", "Выбрать колонки")
	flag.StringVar(&flags.delimiter, "d", "/t", "Использовать другой разделитель ")
	flag.BoolVar(&flags.separated, "s", false, "Только строки с разделителем")

	flag.Parse()

	if flags.fields != "" {
		arr := strings.Split(flags.fields, "/")
		columns := make([]int, len(arr))
		var err error
		for i := 0; i < len(columns); i++ {
			columns[i], err = strconv.Atoi(arr[i])
			if err != nil {
				return nil, err
			}
		}
		sort.Ints(columns)
		flags.setColumns = columns
	}

	return &flags, nil
}
