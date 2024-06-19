package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	url := "https://httpbin.org/get"
	parallelRequest := 5
	requestCount := 50
	result := benchRequest(url, parallelRequest, requestCount)
	for i := 0; i < requestCount; i++ {
		statusCode := <-result
		if statusCode != 200 {
			panic(fmt.Sprintf("Ошибка при отправке запроса: %d", statusCode))
		}
	}
	fmt.Println("Все горутины завершили работу")
}

func benchRequest(url string, parallelRequest int, requestCount int) <-chan int {
	var wg sync.WaitGroup
	var result = make(chan int, requestCount)
	var semaphore = make(chan int, parallelRequest)
	for i := 0; i < requestCount; i++ {
		wg.Add(1)
		semaphore <- i
		go func(semaphore chan int) {
			defer wg.Done()
			statusCode, err := httpRequest(url)
			if err != nil {
				log.Fatal(err)
			}
			result <- statusCode
			<-semaphore
		}(semaphore)
	}
	go func() {
		wg.Wait()
		close(result)
		close(semaphore)
	}()
	return result
}

func httpRequest(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
