package fib

import (
	"testing"
)

// Fib performance comparison function
func BenchmarkFib(b *testing.B) {
	b.ResetTimer() // clean time
	for i := 0; i < b.N; i++ {
		Fib(2)
	}
}

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}
func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}
func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}

// func BenchmarkFib50(b *testing.B) {
// 	benchmarkFib(b, 50)
// }

// go test -bench=. -benchmem
