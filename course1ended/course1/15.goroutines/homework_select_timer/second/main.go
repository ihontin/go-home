package main

// func main() {
// timeNowIs := time.Now()
// AFP := make(chan string)
//
//	go func() {
//		for i := 0; i < 5; i++ {
//			time.Sleep(time.Millisecond * 1200) // выводит сообщение каждые 1200 мс
//			AFP <- fmt.Sprintf("news #%d", i)
//		}
//	}()
//
//	for alive := true; alive; {
//		timer := time.NewTimer(time.Second * 4)
//		select {
//		case news := <-AFP: // выводит сообщение каждые 1200 мс
//			timer.Stop()
//			fmt.Println(news)
//		case <-timer.C:
//			alive = false
//			fmt.Println("No news in an hour. Service aborting.")
//		case <-time.After(1 * time.Second):
//			alive = false
//			fmt.Println("Timeout reached. Service aborting.")
//		}
//	}
//
// fmt.Println(time.Since(timeNowIs))
// }
// ---------------------------------------------1
//func main() {
//	go printNumbers()
//	// Ждем некоторое время, чтобы горутина успела выполниться
//	time.Sleep(5 * time.Second)
//
//}
//func printNumbers() {
//	for i := 1; i <= 10; i++ {
//		fmt.Println(i)
//		time.Sleep(time.Second)
//	}
//}

// ---------------------------------------------2
//func printID(wg *sync.WaitGroup, id int) {
//	defer wg.Done()
//	fmt.Println("Горутина", id)
//}
//func main() {
//	runtime.GOMAXPROCS(2) // Устанавливаем количество процессоров для параллельного выполнения
//
//	var wg sync.WaitGroup
//	wg.Add(2)
//
//	go printID(&wg, 1)
//	go printID(&wg, 2)
//
//	wg.Wait()
//}
//------------------------------------------------3

//func sendData(ch chan<- int) {
//	for i := 1; i <= 5; i++ {
//		ch <- i
//	}
//	close(ch)
//}
//
//func receiveData(ch <-chan int) {
//	for num := range ch {
//		fmt.Println("Получено число:", num)
//	}
//}
//
//func main() {
//	ch := make(chan int)
//
//	go sendData(ch)
//	receiveData(ch)
//}
