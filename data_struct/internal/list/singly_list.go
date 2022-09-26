// src: https://dev.to/divshekhar/golang-linked-list-data-structure-h20
package list

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	len  int
	head *Node
	tail *Node
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l *LinkedList) Head() *Node {
	node := l.head
	return node
}

func (l *LinkedList) Tail() *Node {
	node := l.tail
	return node
}

func (l *LinkedList) Display() {
	cur := l.head
	for cur != nil {
		if cur.next != nil {
			fmt.Printf("%v -> ", cur.data)
		} else {
			fmt.Printf("%v", cur.data)
		}
		cur = cur.next
	}
	fmt.Println()
}

func (l *LinkedList) PushFront(data int) {
	el := &Node{
		data: data,
	}

	if l.head == nil {
		l.head = el
		l.tail = el
		l.len++
	} else {
		el.next = l.head
		l.head = el
	}
}

func (l *LinkedList) PushBack(data int) {
	el := &Node{
		data: data,
	}
	if l.head == nil {
		l.head = el
		l.tail = el
		l.len++
	} else {
		l.tail.next = el
		l.tail = el
		l.len++
	}
}

func (l *LinkedList) PopBack() (*Node, error) {
	if l.len == 0 {
		return nil, fmt.Errorf("list is empty")
	}

	if l.len == 1 {
		l.head = nil
		l.tail = nil
		l.len = 0
		return l.head, nil
	}

	cur := l.head
	for ; cur.next != l.tail; cur = cur.next {
	}

	cur.next = nil
	l.tail = cur
	l.len--
	return cur, nil
}

func (l *LinkedList) PopFront() (*Node, error) {
	if l.len == 0 {
		return nil, fmt.Errorf("list is empty")
	}

	if l.len == 1 {
		l.head = nil
		l.tail = nil
		l.len = 0
		return l.head, nil
	}

	cur := l.head
	l.head = l.head.next

	return cur, nil
}
