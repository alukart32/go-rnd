package chainofresponsibility

import "fmt"

type Handler interface {
	Action(eventType int, msg string) string
}

type ConcreteHandlerA struct {
	next Handler
}

func (h ConcreteHandlerA) Action(eventType int, msg string) string {
	if eventType == 1 {
		return fmt.Sprint("concrete handler A in action: ", msg)
	} else if h.next != nil {
		return h.next.Action(eventType, msg)
	}
	return ""
}

type ConcreteHandlerB struct {
	next Handler
}

func (h ConcreteHandlerB) Action(eventType int, msg string) string {
	if eventType == 2 {
		return fmt.Sprint("concrete handler B in action: ", msg)
	} else if h.next != nil {
		return h.next.Action(eventType, msg)
	}
	return ""
}

type ConcreteHandlerC struct {
	next Handler
}

func (h ConcreteHandlerC) Action(eventType int, msg string) string {
	if eventType == 3 {
		return fmt.Sprint("concrete handler C in action: ", msg)
	} else if h.next != nil {
		return h.next.Action(eventType, msg)
	}
	return ""
}
