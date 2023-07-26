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
	file, err := OpenFile(flag.Arg(1))
	if err != nil {
		fmt.Printf("Error opening the file: %v", err)
		os.Exit(1)
	}
	contentsFile := ReadFile(file)
	pattern := flag.Arg(1)
	// Создание паттерна
	re := SetPattern(pattern, flags.ignore)

	searchData := NewSearchData(contentsFile, pattern, flags, re)
	// Фильтрация
	searchData.SearchString()
}
