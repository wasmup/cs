package main

import "testing"

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fib(uint(i))
	}
}
func BenchmarkMemoization(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = memoization(uint(i))
	}
}
func BenchmarkRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = recursion(uint(i))
	}
}
