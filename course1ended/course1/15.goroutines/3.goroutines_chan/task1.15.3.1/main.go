package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	ID       int
	Complete bool
}

var orders []*Order
var completeOrders map[int]bool
var wg sync.WaitGroup
var processTimes chan time.Duration
var sinceProgramStarted time.Duration
var count int
var limitCount int

func main() {
	count = 30
	limitCount = 5
	processTimes = make(chan time.Duration, count) // буф канал на 30 значений
	orders = GenerateOrders(count)
	completeOrders = GenerateCompleteOrders(count)
	programStart := time.Now()
	LimitSpawnOrderProcessing(limitCount)
	wg.Wait()
	sinceProgramStarted = time.Since(programStart)
	go func() {
		time.Sleep(1 * time.Second)
		close(processTimes)
	}()
	checkTimeDifference(limitCount)
}

func checkTimeDifference(limitCount int) {
	// do not edit
	var averageTime time.Duration
	var orderProcessTotalTime time.Duration
	var orderProcessedCount int
	for v := range processTimes {
		orderProcessedCount++
		orderProcessTotalTime += v
	}
	if orderProcessedCount != count {
		panic("orderProcessedCount != count")
	}
	averageTime = orderProcessTotalTime / time.Duration(orderProcessedCount)
	println("orderProcessTotalTime", orderProcessTotalTime/time.Second)
	println("averageTime", averageTime/time.Second)
	println("sinceProgramStarted", sinceProgramStarted/time.Second)
	println("sinceProgramStarted average", sinceProgramStarted/(time.Duration(orderProcessedCount)*time.Second))
	println("orderProcessTotalTime - sinceProgramStarted", (orderProcessTotalTime-sinceProgramStarted)/time.Second)
	if (orderProcessTotalTime/time.Duration(limitCount)-sinceProgramStarted)/time.Second > 0 {
		panic("(orderProcessTotalTime-sinceProgramStarted)/time.Second > 0")
	}
	fmt.Println(orderProcessTotalTime, time.Duration(limitCount), sinceProgramStarted)
}

func LimitSpawnOrderProcessing(limitCount int) {
	limit := make(chan struct{}, limitCount)
	var t time.Time
	for _, order := range orders {
		wg.Add(1)
		limit <- struct{}{}
		t = time.Now()
		go OrderProcessing(order, limit, t)
	}
	// limit spawn OrderProcessing worker by variable limit
}

func OrderProcessing(order *Order, limit chan struct{}, t time.Time) {
	// complete orders if they completed
	if completeOrders[order.ID] {
		order.Complete = true
	}
	time.Sleep(1 * time.Second)
	processTimes <- time.Since(t)
	wg.Done()
	<-limit
}

func GenerateOrders(count int) []*Order {
	for i := 0; i < count; i++ {
		var un = Order{
			i, false,
		}
		orders = append(orders, &un)
	}
	return orders
	// generate uncomplete orders by count variable
}

func GenerateCompleteOrders(maxOrderID int) map[int]bool {
	//rand.Seed(time.Now().UnixNano())
	var un = make(map[int]bool, 40)
	for i := 0; i < maxOrderID; i++ {
		if i%2 == 0 {
			un[i] = true
		} else {
			un[i] = false
		}
		//r := rand.Intn(100)
		//if r < 50 {
		//	un[i] = true
		//} else {
		//	un[i] = false
		//}
	}
	return un
	// chance 50% to generate map of complete order
}
