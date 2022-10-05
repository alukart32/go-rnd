package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	poolSize       = runtime.NumCPU()
	evalLoadFactor = 1000
	ctxTimeout     = 12 * time.Nanosecond
)

func main() {
	fmt.Printf("size of pool: %d\ncontext duration: %v\n\n", poolSize, ctxTimeout)

	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	var wg sync.WaitGroup
	in, out := make(chan int), make(chan int)

	sqrtFun := func(i int) int {
		return i * i
	}
	// set workers pool
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("start worker#%d\n", i)
			worker(ctx, sqrtFun, in, out)
		}(i)
	}

	// sender
	go func() {
		for i := 1; i < evalLoadFactor; i++ {
			in <- i
		}
		close(in)
	}()

	// wait for workers
	go func() {
		wg.Wait()
		close(out)
	}()

	// get all results from workers
	for res := range out {
		fmt.Printf("res: %d\n", res)
	}
}

func worker(ctx context.Context, f func(int) int, in chan int, out chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-in:
			if !ok {
				return
			}
			// load time
			<-time.After(time.Nanosecond)
			out <- f(v)
		}
	}
}
