package main

import (
	"bufio"
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

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			return
		}
		_, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("ошибка отправки сообщения", err)
			os.Exit(1)
		}
	}
	if err1 := scanner.Err(); err1 != nil {
		fmt.Println("ошибка сканера сообщения", err1)
	}
	fmt.Println("sendToserver is done")

}

func sendToserver(conn net.Conn, t string) {
	if t == "exit" {
		return
	}
	_, err := conn.Write([]byte(t))
	if err != nil {
		fmt.Println("ошибка отправки сообщения", err)
		os.Exit(1)
	}
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
