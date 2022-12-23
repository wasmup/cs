package main

import (
	"fmt"
	"time"
)

func main() {
	for _, f := range []F{fib, memoization, recursion} {
		t0 := time.Now()
		fmt.Println("\n\t fib 42 is", f(42))
		fmt.Println(time.Since(t0))
	}
}

func fib(n uint) (a uint) {
	for b := uint(1); n > 0; n-- {
		a, b = b, a+b
	}
	return a
}

func recursion(n uint) uint {
	if n < 2 {
		return n
	}
	return recursion(n-1) + recursion(n-2)
}

func memoization(n uint) uint {
	if v, ok := m[n]; ok {
		return v
	}
	v, ok := m[n-1]
	if !ok {
		v = memoization(n - 1)
	}
	v += m[n-2]
	m[n] = v
	return v
}

var m = map[uint]uint{0: 0, 1: 1}

type F func(uint) uint
