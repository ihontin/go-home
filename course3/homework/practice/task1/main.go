package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	request, _ := reader.ReadString('\n')
	fmt.Println(request)
	p := strings.Split(request, " ")
	method := p[0]
	path := p[1]
	if method == "GET" && path == "/" {
		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n\r\nHello, World!")
	} else {
		fmt.Fprintf(conn, "HTTP/1.1 404 Not Found\r\n\r\n")
	}
}

func main() {
	// listen port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	for {
		//accept incoming massages
		conn, err1 := ln.Accept()
		if err1 != nil {
			fmt.Println(err1)
			continue
		}
		go handleConnection(conn)
	}
}
