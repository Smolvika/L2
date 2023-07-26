package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Создание флага
	timeout := flag.Int("timeout", 10, "Таймаут на подключение к серверу")
	flag.Parse()

	if len(flag.Args()) < 2 {
		log.Fatal("Check the host and port")
	}
	// Получение хоста и порта
	host := flag.Arg(0)
	port := flag.Arg(1)
	address := fmt.Sprintf("%s:%s", host, port)
	go func() {
		if err := run(address, *timeout); err != nil {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-quit
}

func run(address string, timeout int) error {
	// Подключение к хосту по протоколу TCP
	conn, err := net.DialTimeout("tcp", address, time.Duration(timeout)*time.Second)
	defer conn.Close()

	if err != nil {
		return err
	}
	// Обработка сообщений
	for {
		var msg string
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			msg += scanner.Text() + "\n"
		}
		if _, err = conn.Write([]byte(msg)); err != nil {
			return err
		}

		buf := make([]byte, 1024)

		n, err := conn.Read(buf)

		if err != nil {
			return err
		}

		fmt.Println(string(buf[:n]))
		return nil
	}

}
