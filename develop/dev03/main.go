package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Установака флагов
	flags := ParseFlags()
	// Открытие и чтение файла
	data, err := OpenAndReadFile(flags.input)
	if err != nil {
		fmt.Printf("error working with the file: %v", err)
		os.Exit(1)
	}
	// Удаление дубликатов
	if flags.u {
		data = RemoveDuplicate(data)
	}
	// Сортировка строк
	Sort(data, flags)

	if err := ChangedResult(flag.Arg(0), data); err != nil {
		fmt.Printf("Error when changing the file %s: %v", flag.Arg(0), err)
	}
}
