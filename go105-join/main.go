package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, t0 time.Time) {
			log.Println("WaitGroup start:", i, time.Since(t0))
			time.Sleep(1000 * time.Millisecond) // a job todo
			log.Println("WaitGroup stop :", i, time.Since(t0))
			wg.Done()
		}(i, time.Now())
	}
	////////////////////////////////////////////////////////////////////////////
	ch := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		go func(i int, t0 time.Time) {
			ch <- struct{}{} // counting semaphore
			log.Println("chan start:", i, time.Since(t0))
			time.Sleep(100 * time.Millisecond) // a job todo
			log.Println("chan stop :", i, time.Since(t0))
			<-ch // free
		}(i, time.Now())
	}
	////////////////////////////////////////////////////////////////////////////
	var mu sync.Mutex
	mu.Lock()
	go func(t0 time.Time) {
		log.Println("Mutex start:", time.Since(t0))
		time.Sleep(1000 * time.Millisecond) // a job todo
		log.Println("Mutex stop :", time.Since(t0))
		mu.Unlock()
	}(time.Now())

	t0 := time.Now()
	mu.Lock() // join
	fmt.Println("Mutex finished", time.Since(t0))

	wg.Wait() // join
	fmt.Println("WaitGroup finished")
}
