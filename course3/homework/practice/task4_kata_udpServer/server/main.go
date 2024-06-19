// Код сервера
package main

import (
	"fmt"
	"net"
)

func main() {
	// Создаем адрес для прослушивания
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Ошибка при создании адреса:", err)
		return
	}

	// Создаем соединение
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Сервер запущен. Ожидание сообщений...")

	// Принимаем входящие сообщения
	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Ошибка при чтении сообщения:", err)
			continue
		}

		// Вывод полученного сообщения
		fmt.Printf("Получено сообщение от клиента %s: %s\n", addr.String(), string(buffer[:n]))

		// Отправка ответа клиенту
		response := "Привет от сервера!"
		_, err = conn.WriteToUDP([]byte(response), addr)
		if err != nil {
			fmt.Println("Ошибка при отправке ответа:", err)
			continue
		}

		fmt.Println("Ответ успешно отправлен клиенту.")
	}
}
