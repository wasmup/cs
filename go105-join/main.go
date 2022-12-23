package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func async(i int, t0 time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println(i, time.Since(t0))
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go async(i, time.Now(), &wg)
	}

	wg.Wait() // join
	fmt.Println("done")
	////////////////////////////////////////////////////////////////////////////
	ch := make(chan int, 10)
	for i := 0; i < cap(ch); i++ {
		go fch(i, ch, time.Now())
	}

	sum := 0
	for i := 0; i < cap(ch); i++ {
		sum += <-ch // join
	}
	fmt.Println("sum =", sum)
	////////////////////////////////////////////////////////////////////////////
	var mu sync.Mutex
	mu.Lock() // join
	go func(t0 time.Time) {
		defer mu.Unlock()
		log.Println("long running job", time.Since(t0))
		time.Sleep(1000 * time.Millisecond)
	}(time.Now())

	t0 := time.Now()
	mu.Lock()
	fmt.Println("finished", time.Since(t0))
}

func fch(i int, ch chan int, t0 time.Time) {
	log.Println(i, time.Since(t0))
	ch <- i
}
