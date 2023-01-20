package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	log.SetFlags(0)

	const n = 17
	var addressed int
	var data = make(chan int, 1)
	var barrier, all sync.WaitGroup

	for i := 0; i < n; i++ {
		barrier.Add(1) // A
		for i := 0; i < n; i++ {
			all.Add(1)
			go fun(i, &addressed, data, &barrier, &all)
		}

		addressed = i
		data <- i * i
		barrier.Done() // B

		all.Wait() // C
	}
}

func fun(i int, addressed *int, data chan int, barrier, all *sync.WaitGroup) {
	defer all.Done()     // C
	barrier.Wait()       // A
	if *addressed != i { // B
		return
	}
	v := <-data
	log.Println(i, v)
	time.Sleep(100 * time.Millisecond) // e.g. a long running task
}
