package observer

import (
	"fmt"
	"testing"
)

type TestObserver struct {
	ID      string
	Message string
}

func (t *TestObserver) Notify(m string) {
	fmt.Printf("Observer %s: message '%s' received \n", t.ID, m)
	t.Message = m
}

func newTestObserver(id string, m string) *TestObserver {
	return &TestObserver{
		ID:      id,
		Message: m,
	}
}

func TestPublisher(t *testing.T) {
	observer1 := newTestObserver("1", "")
	observer2 := newTestObserver("2", "")
	observer3 := newTestObserver("3", "")
	publisher := &Publisher{}

	t.Run("AddObserver", func(t *testing.T) {
		publisher.AddObserver(observer1)
		publisher.AddObserver(observer2)
		publisher.AddObserver(observer3)

		if len(publisher.Observers) != 3 {
			t.Fail()
		}
	})

	t.Run("RemoveObserver", func(t *testing.T) {
		publisher.RemoveObserver(observer1)

		if len(publisher.Observers) != 2 {
			t.Errorf("The size of the observer list is not the "+
				"expected. 3 != %d\n", len(publisher.Observers))
		}

		for _, o := range publisher.Observers {
			testObserver, ok := o.(*TestObserver)
			if !ok {
				t.Fail()
			}

			if testObserver.ID == "1" {
				t.Fail()
			}
		}
	})

	t.Run("Notify", func(t *testing.T) {
		for _, observer := range publisher.Observers {
			testObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
				break
			}

			if testObserver.Message != "" {
				t.Errorf("The observer's Message field weren't "+"  empty: %s\n", testObserver.Message)
			}
		}
		message := "Hello World!"
		publisher.NotifyObservers(message)

		for _, observer := range publisher.Observers {
			printObserver, ok := observer.(*TestObserver)
			if !ok {
				t.Fail()
				break
			}

			if printObserver.Message != message {
				t.Errorf("expected message on observer %s was "+
					"not expected: '%s' != '%s'\n", printObserver.ID,
					printObserver.Message, message)
			}
		}
	})
}
