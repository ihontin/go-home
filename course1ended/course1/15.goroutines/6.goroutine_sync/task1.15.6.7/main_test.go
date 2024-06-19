package main

import (
	"testing"
)

func BenchmarkWithoutPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		p = &Person{i}
		_ = p.Age
	}
}

func BenchmarkWithPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		per := personPool.Get().(*Person)
		per.Age = i
		p = per
		personPool.Put(per)
		_ = p.Age
	}
}

//go test -bench=. -benchmem -benchtime=1000x
//BenchmarkWithoutPool-8   	    1000	        23.60 ns/op	       8 B/op	       1 allocs/op
//BenchmarkWithPool-8      	    1000	        12.33 ns/op	       1 B/op	       0 allocs/op
