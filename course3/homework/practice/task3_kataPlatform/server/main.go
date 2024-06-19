package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Server is running")

	for {
		fmt.Println("loop on the server")
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println(err1)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err1 := conn.Read(buf)
	if err1 != nil {
		log.Fatal(err1)
		return
	}

	fmt.Println("Data received from client:", string(buf[:n]))

	resp := "Response from Server"
	_, err := conn.Write([]byte(resp))
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Server response sent")

}
