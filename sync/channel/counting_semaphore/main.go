package main

import (
	"log"
	"math/rand"
	"time"
)

type Seat int
type Bar chan Seat

func (b Bar) ServeCustomerAt(id int, seat Seat) {
	log.Print("++ customer#", id, " drinks at seat#", seat)
	time.Sleep(time.Second * time.Duration(3+rand.Intn(16)))
	log.Print("-- customer#", id, " frees seat#", seat)
	b <- seat // free seat and leave the bar
}

func main() {
	customerGoToTheBar()
}

func customerWaitSeat() {
	rand.Seed(time.Now().UnixNano())

	// the bar with 10 seats
	bar := make(Bar, 10)
	// set seats
	for id := 0; id < cap(bar); id++ {
		bar <- Seat(id)
	}

	// serve customers that get seats
	for customerId := 0; ; customerId++ {
		seat := <-bar
		go bar.ServeCustomerAt(customerId, seat)
	}

	for {
		time.Sleep(time.Second)
	}
}

type Customer struct{ id int }
type Bar2 chan Customer

func (bar Bar2) ServeCustomer(c Customer) {
	log.Print("++ customer#", c.id, " starts drinking")
	time.Sleep(time.Second * time.Duration(3+rand.Intn(16)))
	log.Print("-- customer#", c.id, " leaves the bar")
	<-bar // leaves the bar and save a space
}

func customerGoToTheBar() {
	rand.Seed(time.Now().UnixNano())

	// The bar can serve most 10 customers
	// at the same time.
	bar24x7 := make(Bar2, 10)
	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second * 2)
		customer := Customer{customerId}
		// Wait to enter the bar.
		bar24x7 <- customer
		go bar24x7.ServeCustomer(customer)
	}
	for {
		time.Sleep(time.Second)
	}
}
