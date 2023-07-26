package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Создание флага
	fileFlag := flag.String("-O", "index.html", "Файл в который будут сохранены полученные данные")
	flag.Parse()
	// Выполнение запрос
	resp, err := http.Get(flag.Arg(0))
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Создание файл для записи
	file, err := os.Create(*fileFlag)
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// Копирование данных в файл
	_, err = io.Copy(file, resp.Body)
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

}
