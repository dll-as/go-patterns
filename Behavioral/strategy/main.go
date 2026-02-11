package main

import "fmt"

// Strategy interface
type PaymentStrategy interface {
	Pay(amount float64)
}

// Concrete strategies
type CreditCardPayment struct{}

func (c *CreditCardPayment) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Credit Card\n", amount)
}

type PayPalPayment struct{}

func (p *PayPalPayment) Pay(amount float64) {
	fmt.Printf("Paid %.2f using PayPal\n", amount)
}

// Context
type PaymentContext struct {
	strategy PaymentStrategy
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) Pay(amount float64) {
	p.strategy.Pay(amount)
}

func main() {
	payment := &PaymentContext{}

	payment.SetStrategy(&CreditCardPayment{})
	payment.Pay(100.50) // Paid 100.50 using Credit Card

	payment.SetStrategy(&PayPalPayment{})
	payment.Pay(200.75) // Paid 200.75 using PayPal
}
