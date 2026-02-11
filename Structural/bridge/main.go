package main

import "fmt"

// Implementor (Platform)
type Platform interface {
	Send(message string)
}

type EmailPlatform struct{}

func (e *EmailPlatform) Send(message string) {
	fmt.Println("Email:", message)
}

type SMSPlatform struct{}

func (s *SMSPlatform) Send(message string) {
	fmt.Println("SMS:", message)
}

// Abstraction (Notification)
type Notification struct {
	platform Platform
}

func (n *Notification) Send(message string) {
	n.platform.Send(message)
}

// Refined Abstraction (specific types)
type MessageNotification struct {
	Notification
}

type AlertNotification struct {
	Notification
}

func main() {
	emailPlatform := &EmailPlatform{}
	smsPlatform := &SMSPlatform{}

	message := &MessageNotification{Notification{platform: emailPlatform}}
	alert := &AlertNotification{Notification{platform: smsPlatform}}

	message.Send("Hello via Email!")
	alert.Send("Alert via SMS!")
}
