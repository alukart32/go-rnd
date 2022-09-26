package main

import (
	"context"
	"fmt"
	"log"
	"sync"
)

func main() {
	ctx := context.Background()
	computingRunsSetOfGoroutines(ctx)
}

type CtxChannelName string

var (
	wg    sync.WaitGroup
	tasks = 5
)

func computingRunsSetOfGoroutines(ctx context.Context) error {
	ctx, cancelCtx := context.WithCancel(ctx)
	c := make(chan int)
	wg.Add(tasks + 1)

	go compute(ctx, c)
	for i := 0; i < tasks; i++ {
		c <- i
	}
	cancelCtx()
	wg.Wait()
	close(c)

	fmt.Printf("computing is finished\n")
	return nil
}

func compute(ctx context.Context, rch <-chan int) error {
	for {
		// blocking until one of case receive or send operation is invoked
		select {
		case <-ctx.Done():
			if ctx.Err() != nil {
				log.Printf("compute err: %s", ctx.Err())
			}
			wg.Done()
			return nil
		case v := <-rch:
			log.Printf("receive from channel: %v\n", v)
			wg.Done()
		}
	}
}
