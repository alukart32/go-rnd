// https://go101.org/article/channel-closing.html
// 4. A variant of the "M receivers, one sender" situation: the close request is made by a third-party goroutine
// why to use stop func
package main

import (
	"fmt"
	"math/rand"

	"time"
)

func main() {
	p := NewProducer()
	// receiver
	go func() {
		for n := range p.Stream() {
			fmt.Println(n)
			time.Sleep(time.Millisecond)
		}
	}()
	time.Sleep(100 * time.Millisecond)
	p.Close()
	fmt.Println("closed")
}

// sender
type Producer struct {
	data chan int

	// closing is sent a value when the producer is requested
	// to close down.
	closing chan struct{}
	// closed is closed when the producer has closed down.
	closed chan struct{}
}

func NewProducer() *Producer {
	p := &Producer{
		data:    make(chan int),
		closing: make(chan struct{}),
		closed:  make(chan struct{}),
	}
	go p.run()
	return p
}

func (p *Producer) Stream() chan int {
	return p.data
}

func (p *Producer) run() {
	for {
		select {
		case <-p.closing:
			close(p.data)
			close(p.closed)
			return
		case p.data <- rand.Int():
		}
	}
}

func (p *Producer) Close() {
	select {
	case p.closing <- struct{}{}:
		// Wait for the run goroutine to clean itself up.
		<-p.closed
	case <-p.closed:
		// The producer has already closed down.
	}
}
