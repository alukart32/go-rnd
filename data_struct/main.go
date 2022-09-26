package main

import (
	"fmt"

	"alukart32.com/usage/datastruct/internal/list"
)

func main() {
	list := &list.LinkedList{}

	// list.PushFront(1)
	// list.PushFront(2)
	// list.PushFront(3)
	// list.PushFront(4)

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)

	list.Display()
	fmt.Printf("List size: %d\n", list.Len())
}
