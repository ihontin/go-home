package main

import (
	"fmt"
	"time"
)

func main() {
	//AFP := make(chan string)
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		time.Sleep(time.Millisecond * 200)
	//		AFP <- fmt.Sprintf("news #%d", i)
	//	}
	//}()
	//for {
	//	select {
	//	case news := <-AFP:
	//		fmt.Println(news)
	//	case <-time.After(time.Second):
	//		fmt.Println("No news in an hour.")
	//		return
	//	}
	//}

	//-------------------------------------------1
	//channel1 := make(chan int)
	//channel2 := make(chan string)
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	channel1 <- 42
	//}()
	//go func() {
	//	time.Sleep(3 * time.Second)
	//	channel2 <- "Hello, Golang!"
	//}()
	//select {
	//case value := <-channel1:
	//	fmt.Println("Получено значение из channel1:", value)
	//case message := <-channel2:
	//	fmt.Println("Получено сообщение из channel2:", message)
	//}

	//--------------------------------------------- 2

	//channel1 := make(chan int)
	//channel2 := make(chan string)
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	channel1 <- 42
	//}()
	//go func() {
	//	time.Sleep(3 * time.Second)
	//	channel2 <- "Hello, Golang!"
	//}()
	//select {
	//case value := <-channel1:
	//	fmt.Println("Получено значение из channel1:", value)
	//case message := <-channel2:
	//	fmt.Println("Получено сообщение из channel2:", message)
	//default:
	//	fmt.Println("Ни один из каналов не готов")
	//}

	//-------------------------------------------- 3

	channel := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello, Golang!"
	}()
	select {
	case message := <-channel:
		fmt.Println("Получено сообщение из channel:", message)
	case <-time.After(3 * time.Second):
		fmt.Println("Превышено время ожидания")
	}

	//------------------------------------------ 4

	//channel := make(chan int)
	//select {
	//case channel <- 42:
	//	fmt.Println("Значение успешно отправлено в канал")
	//case <-time.After(2 * time.Second):
	//	fmt.Println("Превышено время ожидания")
	//}

}
