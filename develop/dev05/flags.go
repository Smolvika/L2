package main

import "flag"

type Flags struct {
	after   int
	before  int
	context int
	count   bool
	ignore  bool
	invert  bool
	fixed   bool
	num     bool
}

func ParseFlags() *Flags {
	var flags Flags
	flag.IntVar(&flags.after, "A", 0, "Число строк после найденной")
	flag.IntVar(&flags.before, "B", 0, "Число строк до найденной")
	flag.IntVar(&flags.context, "C", 0, "Число строк до и после найденной")
	flag.BoolVar(&flags.count, "c", false, "Количество строк")
	flag.BoolVar(&flags.ignore, "i", false, "Игнорировать регистр")
	flag.BoolVar(&flags.invert, "v", false, "Вместо совпадения, совпадения исключать")
	flag.BoolVar(&flags.fixed, "F", false, "Точное совпадение со строкой")
	flag.BoolVar(&flags.num, "n", false, "Печатать номер строки ")

	flag.Parse()

	return &flags
}
