package main

import "fmt"

func main() {
	mutex := make(chan struct{}, 1)

	v := 0
	eval := func() {
		mutex <- struct{}{}
		v++
		<-mutex
	}

	f := func(done chan<- struct{}) {
		for i := 0; i < 256; i++ {
			eval()
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go f(done)
	go f(done)
	<-done
	<-done
	fmt.Println(v)
}
