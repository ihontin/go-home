// Код сервера
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type client struct {
	conn net.Conn
	name string
	ch   chan<- string
}

var (
	// канал для всех входящих клиентов
	entering = make(chan client)
	// канал для сообщения о выходе клиента
	leaving = make(chan client)
	// канал для всех сообщений
	messages = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

// broadcaster рассылает входящие сообщения всем клиентам
// следит за подключением и отключением клиентов
func broadcaster() {
	// здесь хранятся все подключенные клиенты
	clients := make(map[client]bool)
	for {
		var mes string
		select {
		case mes = <-messages:
			for key := range clients {
				_, err := key.conn.Write([]byte(mes))
				if err != nil {
					fmt.Println("ошибка отправки сообщения клиенту:", key.name, "func broadcaster:", err)
					return
				}
			}
		case newClient := <-entering:
			clients[newClient] = true
		case lived := <-leaving:
			delete(clients, lived)
		}
	}
}

// handleConn обрабатывает входящие сообщения от клиента
func handleConn(conn net.Conn) {
	defer conn.Close()
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	cli := client{conn, who, ch}
	var text string
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	scaner := bufio.NewScanner(conn)
	for scaner.Scan() {
		text = scaner.Text()
		messages <- who + ": " + text
	}
	if err := scaner.Err(); err != nil {
		fmt.Println("ошибка прочтения", err)
	}
	leaving <- cli
	messages <- who + " has left"
}

// clientWriter отправляет сообщения текущему клиенту
func clientWriter(conn net.Conn, ch <-chan string) {
	_, err := conn.Write([]byte(<-ch))
	if err != nil {
		fmt.Println("ошибка clientWriter сообщения текущему клиенту:", err)
		return
	}
}

//Сервер должен принимать входящие соединения на порту 8000.
//Когда клиент подключается, сервер должен отправить сообщение «You are [адрес клиента]»
//Когда клиент отправляет сообщение, сервер должен отправить его всем подключенным клиентам в формате «[адрес клиента]: [сообщение]».
//Когда клиент отключается, сервер должен отправить сообщение «[адрес клиента] has left».
//Клиент должен подключаться к серверу на порту 8000.
//Клиент должен отправлять введенные с клавиатуры сообщения на сервер.
//Клиент должен выводить полученные сообщения от сервера на экран.
