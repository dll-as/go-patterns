package main

import "fmt"

// Template
type Beverage interface {
	PrepareRecipe()
}

type CaffeineBeverage struct{}

func (c *CaffeineBeverage) PrepareRecipe() {
	c.BoilWater()
	c.Brew()
	c.PourInCup()
	c.AddCondiments()
}

func (c *CaffeineBeverage) BoilWater() {
	fmt.Println("Boiling water")
}

func (c *CaffeineBeverage) PourInCup() {
	fmt.Println("Pouring into cup")
}

// Steps to override
func (c *CaffeineBeverage) Brew()          {}
func (c *CaffeineBeverage) AddCondiments() {}

// Subclass: Tea
type Tea struct {
	CaffeineBeverage
}

func (t *Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (t *Tea) AddCondiments() {
	fmt.Println("Adding lemon")
}

// Subclass: Coffee
type Coffee struct {
	CaffeineBeverage
}

func (c *Coffee) Brew() {
	fmt.Println("Dripping coffee through filter")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("Adding sugar and milk")
}

func main() {
	tea := &Tea{}
	coffee := &Coffee{}

	fmt.Println("Making Tea:")
	tea.PrepareRecipe()

	fmt.Println("\nMaking Coffee:")
	coffee.PrepareRecipe()
}
