package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Println("ошибка создания адреса", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("ошибка подключения к серверу", err)
		return
	}

	message := "Привет от клиента!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("ошибка отправки сообщения", err)
		return
	}
	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("ошибка чтения ответа от сервера", err)
		return
	}
	fmt.Println("Ответ от сервера:", string(buffer[:n]))
}
