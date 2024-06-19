package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("ошибка при прослушивании сервера Listen:", err)
	}
	defer listener.Close()

	for {
		conn, err1 := listener.Accept()
		if err1 != nil {
			fmt.Println("ошибка принятия данных Accept:", err1)
		}
		go hendleConnection(conn)
	}
}

func hendleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	req, _ := reader.ReadString('\n')
	parts := strings.Split(req, " ")
	method := parts[0]
	path := parts[1]

	if method == "GET" && path == "/" {
		_, pa, _, _ := runtime.Caller(0)
		indexFile := filepath.Dir(filepath.Dir(filepath.Dir(pa))) + "/static/index.html"
		file, err1 := os.Open(indexFile)
		if err1 != nil {
			fmt.Println("ошибка открытия файла Open:", err1)
			return
		}
		defer file.Close()
		html, err := io.ReadAll(file)
		if err != nil {
			fmt.Println("ошибка чтения файла ReadAll:", err)
			return
		}
		m := "HTTP/1.1 200 OK\r\n"
		m += "Context-Type: text/html\r\n"
		m += "\r\n"
		m += string(html)
		_, _ = conn.Write([]byte(m))
	} else {
		fmt.Fprintf(conn, "HTTP/1.1 404 Not Found\r\n\r\n")
	}

}
