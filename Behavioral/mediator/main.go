package main

import "fmt"

// Mediator interface
type ChatRoom interface {
	SendMessage(sender, message string)
}

// Concrete Mediator
type ChatRoomImpl struct{}

func (c *ChatRoomImpl) SendMessage(sender, message string) {
	fmt.Printf("[%s] %s\n", sender, message)
}

// Colleague
type User struct {
	name string
	room ChatRoom
}

func (u *User) Send(message string) {
	u.room.SendMessage(u.name, message)
}

func main() {
	room := &ChatRoomImpl{}

	alice := &User{name: "Alice", room: room}
	bob := &User{name: "Bob", room: room}

	alice.Send("Hello Bob!")
	bob.Send("Hi Alice!")
}
