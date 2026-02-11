package main

import "fmt"

// Target interface
type Payment interface {
	Pay(amount float64)
}

// Adaptee (external library)
type Stripe struct{}

func (s *Stripe) MakePayment(amount float64) {
	fmt.Printf("Stripe: processing payment %.2f$\n", amount)
}

// Adapter
type StripeAdapter struct {
	stripe *Stripe
}

func (a *StripeAdapter) Pay(amount float64) {
	a.stripe.MakePayment(amount)
}

func main() {
	// client code only knows Payment interface
	var payment Payment

	// use Stripe via adapter
	payment = &StripeAdapter{stripe: &Stripe{}}
	payment.Pay(100.50)
}
