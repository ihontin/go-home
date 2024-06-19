package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	message := "Письмо от клиента!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Данные от сервера:", string(buffer[:n]))

}
