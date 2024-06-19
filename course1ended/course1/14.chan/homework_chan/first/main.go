package main

import "fmt"

// ------------------------------------1
//func main() {
//	go func() {
//		fmt.Println("Hello, World!")
//	}()
//	time.Sleep(time.Second) // Ждем 1 секунду, чтобы горутина успела выполниться
//}

//------------------------------------2

// В данном примере показано, как закрыть канал после передачи всех данных.
//func main() {
//	ch := make(chan int)
//	go func() {
//		defer close(ch)
//		for i := 1; i <= 5; i++ {
//			ch <- i
//		}
//	}()
//	for num := range ch {
//		fmt.Println(num)
//	}
//}

//------------------------------------3
//В этом примере показано, как проверить, закрыт ли канал, перед тем как получить данные из него.
//func main() {
//	ch := make(chan int)
//
//	go func() {
//		defer close(ch)
//		for i := 1; i <= 5; i++ {
//			ch <- i
//		}
//	}()
//
//	for {
//		num, ok := <-ch
//		if !ok {
//			break
//		}
//		fmt.Println(num)
//	}
//}

// ------------------------------------4
// В этом примере показано, как объединить данные из нескольких каналов в один.
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	mergedCh := make(chan int)

	go func() {
		defer close(ch1)
		for i := 1; i <= 5; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 6; i <= 10; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer close(mergedCh)
		for num := range ch1 {
			mergedCh <- num
		}
	}()

	go func() {
		for num := range ch2 {
			mergedCh <- num
		}
	}()

	for num := range mergedCh {
		fmt.Println(num)
	}
}

// ------------------------------------5
// Ограничение количества горутин с помощью канала
// В этом примере показано, как использовать канал для ограничения количества одновременно выполняющихся горутин.
//func main() {
//	ch := make(chan struct{}, 3)
//	var wg sync.WaitGroup
//
//	for i := 1; i <= 5; i++ {
//		wg.Add(1)
//		ch <- struct{}{}
//		go func(num int) {
//			defer func() {
//				<-ch
//				wg.Done()
//			}()
//			fmt.Println("Горутина", num, "выполняется")
//		}(i)
//	}
//
//	wg.Wait()
//}

//------------------------------------6

//func longTimeRequest() <-chan int32 {
//	r := make(chan int32)
//	go func() {
//		// Simulate a workload.
//		time.Sleep(time.Second * 3)
//		r <- rand.Int31n(100)
//	}()
//	return r
//}
//func sumSquares(a, b int32) int32 {
//	return a*a + b*b
//}
//func main() {
//	rand.Seed(time.Now().UnixNano())
//	a, b := longTimeRequest(), longTimeRequest()
//	fmt.Println(sumSquares(<-a, <-b))
//}

// --------------------------------------7
//func source(c chan<- int32) {
//	ra, rb := rand.Int31(), rand.Intn(3)+1
//	// Sleep 1s/2s/3s.
//	time.Sleep(time.Duration(rb) * time.Second)
//	c <- ra
//}
//
//func main() {
//	rand.Seed(time.Now().UnixNano())
//
//	startTime := time.Now()
//	// c must be a buffered channel.
//	c := make(chan int32, 5)
//	for i := 0; i < cap(c); i++ {
//		go source(c)
//	}
//	// Only the first response will be used.
//	rnd := <-c
//	fmt.Println(time.Since(startTime))
//	fmt.Println(rnd)
//}
