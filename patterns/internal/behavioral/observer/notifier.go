package observer

import "fmt"

type Observer interface {
	Notify(string)
}

type Publisher struct {
	Observers []Observer
}

func (p *Publisher) AddObserver(o Observer) {
	p.Observers = append(p.Observers, o)
}

func (p *Publisher) RemoveObserver(o Observer) {
	var indexToRemove int

	for i, observer := range p.Observers {
		if observer == o {
			indexToRemove = i
			break
		}
	}

	p.Observers = append(p.Observers[:indexToRemove], p.Observers[indexToRemove+1:]...)
}

func (p *Publisher) NotifyObservers(m string) {
	fmt.Printf("Publisher received message '%s' to notify observers\n", m)
	for _, o := range p.Observers {
		o.Notify(m)
	}
}
