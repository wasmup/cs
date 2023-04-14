package main

import (
	"fmt"
	"testing"
)

func benchmarkIsPalindrome(x int, f func(int) bool, b *testing.B) {
	for n := 0; n < b.N; n++ {
		f(x)
	}
}

func BenchmarkIsPalindrome1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkIsPalindrome(121, isPalindrome1, b)
	}
}

func BenchmarkIsPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkIsPalindrome(121, isPalindrome2, b)
	}
}

func main() {
	fmt.Println("Benchmarking isPalindrome1...")
	testing.Benchmark(BenchmarkIsPalindrome1)

	fmt.Println("Benchmarking isPalindrome2...")
	testing.Benchmark(BenchmarkIsPalindrome2)
}
