package main

import (
	"log"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	a := make([]int, 1*1024*1024)
	for i := range a {
		a[i] = i
	}

	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	b := make([]int, len(a))
	copy(b, a)
	c := make([]int, len(a))
	copy(c, a)
	d := make([]int, len(a))
	copy(d, a)

	t0 := time.Now()
	quicksort(a) // O(n**2)
	log.Println(time.Since(t0))

	t0 = time.Now()
	merge(b) // O(n*log(n))
	log.Println(time.Since(t0))

	t0 = time.Now()
	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })
	log.Println(time.Since(t0))

	t0 = time.Now()
	sort.Ints(d)
	log.Println(time.Since(t0))
}

// O(n**2)
func quicksort(a []int) {
	if len(a) > 1 {
		pivot := partition(a)
		quicksort(a[:pivot+1])
		quicksort(a[pivot+1:])
	}
}
func partition(a []int) int {
	pivot := a[0]
	i, j := 0, len(a)-1
	for i < j {
		for ; a[i] <= pivot && i < len(a)-1; i++ {
		}
		for ; a[j] > pivot && j > 0; j-- {
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}
	a[0] = a[j]
	a[j] = pivot
	return j
}

// O(n*log(n))
func merge(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	m := len(a) / 2
	return func(a, b []int) []int {
		i, j := 0, 0
		result := make([]int, len(a)+len(b))
		for k := range result {
			if j >= len(b) || (i < len(a) && a[i] <= b[j]) {
				result[k] = a[i]
				i++
			} else {
				result[k] = b[j]
				j++
			}
		}
		return result
	}(merge(a[:m]), merge(a[m:]))
}
