package main

import "fmt"

// Command interface
type Command interface {
	Execute()
}

// Receiver
type Light struct{}

func (l *Light) TurnOn() {
	fmt.Println("Light is ON")
}

func (l *Light) TurnOff() {
	fmt.Println("Light is OFF")
}

// Concrete commands
type TurnOnCommand struct {
	light *Light
}

func (c *TurnOnCommand) Execute() {
	c.light.TurnOn()
}

type TurnOffCommand struct {
	light *Light
}

func (c *TurnOffCommand) Execute() {
	c.light.TurnOff()
}

// Invoker
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func main() {
	light := &Light{}

	onCommand := &TurnOnCommand{light: light}
	offCommand := &TurnOffCommand{light: light}

	remote := &RemoteControl{}

	remote.SetCommand(onCommand)
	remote.PressButton() // Light is ON

	remote.SetCommand(offCommand)
	remote.PressButton() // Light is OFF
}
