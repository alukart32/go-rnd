package main

import (
	"context"
	"fmt"
)

type message struct {
	respCh chan<- int
	param  string
	ctx    context.Context
}

func ProcessMessage(work <-chan message) {
	for job := range work {
		select {
		case <-job.ctx.Done():
			continue
		default:
		}

		select {
		case <-job.ctx.Done():
		case job.respCh <- len(job.param):
		}
	}
}

func NewMessage(ctx context.Context, q chan message, param string) {
	r := make(chan int)
	select {
	case <-ctx.Done():
		fmt.Println("Context ended before q could see message")
		return
	case q <- message{
		respCh: r,
		param:  param,
		ctx:    ctx,
	}:
	}

	select {
	case <-ctx.Done():
		fmt.Println("Context ended before q could process message")
	case out := <-r:
		fmt.Printf("The len of %s is %d\n", param, out)
	}
}

func main() {
	q := make(chan message)
	go ProcessMessage(q)
	ctx := context.Background()
	NewMessage(ctx, q, "text")
	NewMessage(ctx, q, "bazar")
}
