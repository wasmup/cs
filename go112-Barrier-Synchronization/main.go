package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	log.SetFlags(0)

	const n = 17
	syncMailbox := make([]int, n)
	var barrier, all sync.WaitGroup

	barrier.Add(1) // A
	for i := 0; i < n; i++ {
		all.Add(1)
		go fun(i, syncMailbox, &barrier, &all)
	}

	for i := 0; i < n; i++ {
		syncMailbox[i] = i // fill
	}

	barrier.Done() // B

	all.Wait() // C

	fmt.Println(syncMailbox)
}

func fun(i int, syncMailbox []int, barrier, all *sync.WaitGroup) {
	defer all.Done()    // C
	barrier.Wait()      // A
	v := syncMailbox[i] // B
	log.Println(i, v)
	syncMailbox[i] = i * v
}
