package main

import (
	"log"
	"time"
)

func Timer(t time.Duration) <-chan struct{} {
	c := make(chan struct{})

	go func() {
		log.Println("sleep for " + t.String())
		time.Sleep(t)
		c <- struct{}{}
	}()

	return c
}

func main() {
	log.Println("start main func")
	<-Timer(time.Second)
	log.Println("after sleep 1")
	<-Timer(time.Second)
	log.Println("after sleep 2")
	log.Println("more efficient way: <-time.After")
	<-time.After(time.Second)
	log.Println("after <-time.After")
}
