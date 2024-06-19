package main

import (
	"fmt"
	"studentgit.kata.academy/Alkolex/go-kata/course2/5.optimization/2.optimization_jsoniter/task2.5.2.1/models"
	"testing"
)

// go test -bench=. -benchmem -benchtime=1000x
func BenchmarkJsonMarsh_Marshal(b *testing.B) {
	dataByte, err := GetData()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	var sonMarsh jsonMarsh
	_ = sonMarsh.Unmarshal(dataByte, &sonMarsh)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sonMarsh.Marshal(sonMarsh)
	}
}

func BenchmarkJsonIter_Marshal(b *testing.B) {
	dataByte, err := GetData()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	var sonIter jsonIter
	_ = sonIter.Unmarshal(dataByte, &sonIter)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sonIter.Marshal(sonIter)
	}
}

func BenchmarkEasyMarshJson_Marshal(b *testing.B) {
	dataByte, err := GetData()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	var sonEasy EasyMarshJson
	var sss models.Whetherer
	_ = sonEasy.Unmarshal(dataByte, &sss)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = sonEasy.Marshal(sss)
	}
}

//-----------------------------------------unmarshal

func BenchmarkJsonMarsh_Unmarshal(b *testing.B) {
	dataByte, err := GetData()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	var sonMarsh jsonMarsh

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sonMarsh.Unmarshal(dataByte, &sonMarsh)
	}
}

func BenchmarkJsonIter_Unmarshal(b *testing.B) {
	dataByte, err := GetData()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	var sonIter jsonIter

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sonIter.Unmarshal(dataByte, &sonIter)
	}
}

func BenchmarkEasyMarshJson_Unmarshal(b *testing.B) {
	dataByte, err := GetData()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	var sonEasy EasyMarshJson
	var sss models.Whetherer

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sonEasy.Unmarshal(dataByte, &sss)
	}

}
