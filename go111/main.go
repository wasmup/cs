package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(bits.OnesCount(42)) // 3
	fmt.Println(OnesCount(42))
}

func OnesCount(n int) (result int) {
	for ; n > 0; n >>= 1 {
		result += n & 1
	}
	return
}
