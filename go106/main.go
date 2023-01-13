package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	var list = map[int]int{}
	defer func(start time.Time) {
		if len(list) > 0 {
			fmt.Println("\nPractice List:")
		}
		for k, v := range list {
			fmt.Println(k, "*", v, "=", k*v)
		}
		fmt.Println("\nThe time elapsed since the beginning:", time.Since(start))
	}(time.Now())

	type Result struct {
		ans int
		err error
		d   time.Duration
	}
	var ch = make(chan Result)
	var t = make(chan time.Time)
	go func() {
		for {
			var ans int
			t0 := <-t
			_, err := fmt.Scanln(&ans)
			ch <- Result{ans, err, time.Since(t0)}
			if err != nil || ans == 0 {
				return
			}
		}
	}()

	for correct, wrong := 0.0, 0.0; ; {
		a := 1 + rand.Intn(9)
		b := 1 + rand.Intn(9)
		c := a * b
		fmt.Printf("\n%d * %d = ", a, b)
		t <- time.Now()

		select {
		case r := <-ch:
			if r.err != nil || r.ans == 0 {
				return
			}
			if r.ans != c {
				wrong++
				fmt.Print("\nYou Loose:")
				list[a] = b
			} else {
				correct++
				fmt.Print("\nYou Won:")
			}
			fmt.Printf(" %d * %d = %d correct=%.0f wrong=%.0f percent=%.2f %v\n", a, b, c, correct, wrong, 100*(correct-wrong/3)/(correct+wrong), r.d)

		case <-time.After(10 * time.Second):
			return
		}
	}
}
