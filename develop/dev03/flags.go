package main

import "flag"

type Flags struct {
	k     int
	input string
	n     bool
	r     bool
	u     bool
}

func ParseFlags() *Flags {
	var flags Flags
	flag.IntVar(&flags.k, "k", 0, "Номер колонки для сортировки")
	flag.BoolVar(&flags.n, "n", false, "Сортировать по численному значению")
	flag.BoolVar(&flags.r, "r", false, "Сортировать в обратном порядке")
	flag.BoolVar(&flags.u, "u", false, "Убрать повторяющиеся строки")

	flag.Parse()

	flags.input = flag.Arg(0)

	return &flags
}
