package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	baseKont()
}

func baseKont() {
	ctx := context.Background() //корневой, родительский контекст
	fmt.Println(ctx)

	toDo := context.TODO() // для тестов, не особо использовать
	fmt.Println(toDo)

	// можно передавать значения в контекст, но лучше не надо
	coxtWithVal := context.WithValue(ctx, "key", "value")
	fmt.Println(coxtWithVal.Value("key"))

	//хорошо использовать контекст, который может сообщить о завершении задачи
	//для этого можно создать контекст который умеет завершаться
	coxtWithCancel, cancel := context.WithCancel(ctx) //контекст который можно отменить функцией cancel
	fmt.Println(coxtWithCancel.Err())
	cancel()                          //при завершении контекста в его поле Err кладется ошибка
	fmt.Println(coxtWithCancel.Err()) // context canceled

	//контекст закрывается в той же горутине, в которой открывался

	// WithDeadline похоже на работу <-time.After
	coxtWithDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	defer cancel()

	// можно посмотреть строку по информации в контексте
	fmt.Println(coxtWithDeadline.Deadline())
	// ошибки еще нет
	fmt.Println(coxtWithDeadline.Err())

	// после запуска по истечению таймера в Done() будет записано значение и сработает канал который заблочил горутину
	fmt.Println(<-coxtWithDeadline.Done()) // выведет пустую структуру

	// работает аналогично предыдущей функции
	coxtWithTimeout, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	fmt.Println(<-coxtWithTimeout.Done()) // если не использоать <- то получим канал без блокировки горутины
}
