package main

import (
	"fmt"
	"os"
	"time"
)

type Ball uint64

func play(player string, table chan Ball) {
	lastBall := Ball(1)
	for {
		ball := <-table
		fmt.Println(player, ball)
		ball += lastBall
		if ball < lastBall {
			os.Exit(0)
		}
		lastBall = ball
		table <- ball
		time.Sleep(time.Second)
	}
}

func main() {
	table := make(chan Ball)
	go func() {
		table <- 1
	}()
	go play("A:", table)
	play("B:", table)
}
