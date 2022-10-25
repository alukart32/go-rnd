package chainofresponsibility

import "testing"

func TestChain(t *testing.T) {
	handler := ConcreteHandlerA{
		next: ConcreteHandlerB{
			next: ConcreteHandlerC{},
		},
	}

	msg := "working"
	expected := "concrete handler C in action: " + msg

	if r := handler.Action(3, msg); r != expected {
		t.Error("expected: concrete handler C in action: working, but get ", r)
	}
}
