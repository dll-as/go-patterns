package main

import "fmt"

// Observer interface
type Subscriber interface {
	Update(news string)
}

// Subject interface
type Publisher interface {
	Register(sub Subscriber)
	Remove(sub Subscriber)
	Notify(news string)
}

// Concrete subject
type NewsPublisher struct {
	subscribers []Subscriber
}

func (n *NewsPublisher) Register(sub Subscriber) {
	n.subscribers = append(n.subscribers, sub)
}

func (n *NewsPublisher) Remove(sub Subscriber) {
	for i, s := range n.subscribers {
		if s == sub {
			n.subscribers = append(n.subscribers[:i], n.subscribers[i+1:]...)
			break
		}
	}
}

func (n *NewsPublisher) Notify(news string) {
	for _, sub := range n.subscribers {
		sub.Update(news)
	}
}

// Concrete observer
type EmailSubscriber struct {
	name string
}

func (e *EmailSubscriber) Update(news string) {
	fmt.Printf("[%s] received news: %s\n", e.name, news)
}

func main() {
	publisher := &NewsPublisher{}

	sub1 := &EmailSubscriber{name: "Alice"}
	sub2 := &EmailSubscriber{name: "Bob"}

	publisher.Register(sub1)
	publisher.Register(sub2)

	publisher.Notify("New Go 1.21 released!")
	publisher.Notify("Design patterns tutorial available!")
}
