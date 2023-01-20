package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	log.SetFlags(0)

	const n = 17
	syncMailbox := make([]int, n)
	var barrier, all sync.WaitGroup

	for i := 0; i < 10; i++ {
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

		log.Println(syncMailbox)
	}
}

func fun(i int, syncMailbox []int, barrier, all *sync.WaitGroup) {
	defer all.Done()    // C
	barrier.Wait()      // A
	v := syncMailbox[i] // B
	log.Println(i, v)
	syncMailbox[i] = i * v
	time.Sleep(100 * time.Millisecond) // e.g. a long running task
}
