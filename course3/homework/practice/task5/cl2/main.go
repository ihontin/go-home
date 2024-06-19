package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// подключиться к серверу
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	// запустить горутину, которая будет читать все сообщения от сервера и выводить их в консоль
	go clientReader(conn)
	// читать сообщения от stdin и отправлять их на сервер

	var line string
	fmt.Scanln(&line)
	fmt.Println("отправка сообщения", line)
	_, err = conn.Write([]byte(line))
	if err != nil {
		fmt.Println("ошибка отправки сообщения", err)
		os.Exit(1)
	}

	fmt.Println("отправка завершена")

}

// clientReader выводит на экран все сообщения от сервера
func clientReader(conn net.Conn) {
	defer conn.Close()
	for {
		var buffer = make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Клиент читает сообщение от сервера:", string(buffer[:n]))
	}
}
