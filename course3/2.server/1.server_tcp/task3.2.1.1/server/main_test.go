package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"testing"
)

func TestHandleConn(t *testing.T) {
	// Создаем тестовый сервер
	ln, _ := net.Listen("tcp", "localhost:8000")
	go func() {
		for {
			conn, _ := ln.Accept()
			handleConn(conn)
		}
	}()
	// Подключаемся к серверу
	conn, _ := net.Dial("tcp", ln.Addr().String())
	// Пишем в соединение
	conn.Write([]byte("test message\n"))
	// Читаем из соединения
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	msg := string(buf[:n])

	// Проверяем, что сообщение было корректно обработано
	if !strings.Contains(msg, "You are 127.0.0.1:") {
		t.Error("Failed to handle connection")
	}
}

func TestBroadcaster(t *testing.T) {
	_, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	conn.Write([]byte("test message"))
	// Запускаем broadcaster()
	go broadcaster()

	// Добавляем клиента
	clie := client{conn, "test client", make(chan string)}
	entering <- clie
	messages <- "test message"
	leaving <- clie
	fmt.Println(clie.conn.RemoteAddr().String())
}

func TestClientWriter(t *testing.T) {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ch := make(chan string)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer conn.Close()
			go clientWriter(conn, ch)
			buf := make([]byte, 1024)
			_, err1 := conn.Read(buf)
			if err1 != nil {
				log.Fatal(err1)
				return
			}

		}
	}()

	conne, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conne.Close()
	ch <- "test message"
	var b = make([]byte, 1024)
	n, _ := conne.Read(b)

	msg := string(b[:n])
	if msg != "test message" {
		t.Error("Failed to write message to client")
	}
}
