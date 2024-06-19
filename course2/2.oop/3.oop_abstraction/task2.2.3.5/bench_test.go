package main

import "testing"

//go test -bench=. -benchmem -benchtime=1000x

func BenchmarkWithHashCRC64(b *testing.B) {
	hashM := &HashMap{}
	WithHashCRC64()(hashM)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hashM.hashFunk("test")
	}
	b.StopTimer()
}
func BenchmarkWithHashCRC32(b *testing.B) {
	hashM := &HashMap{}
	WithHashCRC32()(hashM)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = hashM.hashFunk("test")
	}
	b.StopTimer()
}
func BenchmarkWithHashCRC16(b *testing.B) {
	hashM := &HashMap{}
	WithHashCRC16()(hashM)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = hashM.hashFunk("test")
	}
	b.StopTimer()
}
func BenchmarkWithHashCRC8(b *testing.B) {
	hashM := &HashMap{}
	WithHashCRC8()(hashM)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = hashM.hashFunk("test")
	}
	b.StopTimer()
}
