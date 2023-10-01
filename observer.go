package main

import (
	"fmt"
)

type Observer interface {
	Update(string)
}

type Subject interface {
	Attach(Observer)
	Detach(Observer)
	Notify(string)
}

type ConcreteSubject struct {
	observers []Observer
	state     string
}

func (s *ConcreteSubject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) Detach(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) Notify(newState string) {
	s.state = newState
	for _, observer := range s.observers {
		observer.Update(newState)
	}
}

type ConcreteObserver struct {
	name string
}

func (o *ConcreteObserver) Update(newState string) {
	fmt.Printf("%s received an update: %s\n", o.name, newState)
}

func main() {
	subject := &ConcreteSubject{}

	observer1 := &ConcreteObserver{name: "Observer 1"}
	observer2 := &ConcreteObserver{name: "Observer 2"}
	observer3 := &ConcreteObserver{name: "Observer 3"}

	subject.Attach(observer1)
	subject.Attach(observer2)
	subject.Attach(observer3)

	subject.Notify("New State 1")
	subject.Notify("New State 2")

	subject.Detach(observer2)

	subject.Notify("New State 3")
}
