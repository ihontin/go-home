package main

import "testing"

func TestMergeChan2(t *testing.T) {
	// Создаем два канала
	ch1 := make(chan int)
	ch2 := make(chan int)
	merged := mergeChan2(ch1, ch2)
	expected := []int{1, 2, 3, 4, 5}

	go func() {
		for _, exp := range expected {
			val := <-merged
			if val != exp {
				t.Errorf("Expected %d, but got %d", exp, val)
			}
		}
		// Проверяем, что после получения всех значений из merged, канал закрыт
		_, ok := <-merged
		if ok {
			t.Error("Expected channel to be closed")
		}
	}()

	//// Отправляем значения в каналы ch1 и ch2
	//values1 := []int{1, 2, 3}
	//for _, val := range values1 {
	//	ch1 <- val
	//}
	//
	//values2 := []int{4, 5}
	//for _, val := range values2 {
	//	ch2 <- val
	//}
	//
	//// Закрываем каналы ch1 и ch2
	//close(ch1)
	//close(ch2)
}
