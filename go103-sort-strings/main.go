package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []string{"0", "1", "12", "121", "2", "1000"}
	sort.Strings(a)
	fmt.Println(a)
	// [0 1 1000 12 121 2]
}
