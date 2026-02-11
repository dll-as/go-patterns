package main

import "fmt"

// Notification interface
type Notification interface {
	Send(message string)
}

// EmailNotification concrete implementation
type EmailNotification struct{}

func (e *EmailNotification) Send(message string) {
	fmt.Println("Sending EMAIL:", message)
}

// SMSNotification concrete implementation
type SMSNotification struct{}

func (s *SMSNotification) Send(message string) {
	fmt.Println("Sending SMS:", message)
}

// Factory Method
func NewNotification(notificationType string) Notification {
	switch notificationType {
	case "email":
		return &EmailNotification{}
	case "sms":
		return &SMSNotification{}
	default:
		return nil
	}
}

func main() {
	notification := NewNotification("email")
	notification.Send("Hello from Factory Method")

	notification2 := NewNotification("sms")
	notification2.Send("Hello again")
}
