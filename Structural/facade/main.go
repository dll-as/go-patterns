package main

import "fmt"

// Subsystems
type AuthenticationService struct{}

func (a *AuthenticationService) Login(user string) {
	fmt.Println("User", user, "authenticated")
}

type PaymentService struct{}

func (p *PaymentService) Pay(amount float64) {
	fmt.Println("Payment of", amount, "processed")
}

type NotificationService struct{}

func (n *NotificationService) Notify(message string) {
	fmt.Println("Notification sent:", message)
}

// Facade
type ShopFacade struct {
	auth   *AuthenticationService
	pay    *PaymentService
	notify *NotificationService
}

func NewShopFacade() *ShopFacade {
	return &ShopFacade{
		auth:   &AuthenticationService{},
		pay:    &PaymentService{},
		notify: &NotificationService{},
	}
}

func (s *ShopFacade) BuyProduct(user string, amount float64) {
	s.auth.Login(user)
	s.pay.Pay(amount)
	s.notify.Notify("Purchase completed for " + user)
}

func main() {
	facade := NewShopFacade()
	facade.BuyProduct("Alice", 99.99)
}
