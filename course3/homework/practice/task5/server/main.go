package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	lestener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Listen error:", err)
	}

	for {
		conn, err := lestener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		scaner := bufio.NewScanner(conn)
		fmt.Println("server stopped")
		for scaner.Scan() {
			fmt.Println("what next")
			text := scaner.Text()
			fmt.Println("handleConn for", text)
		}
	}
}
