package main

import "fmt"

// Element interface
type Item interface {
	Accept(visitor Visitor)
}

// Concrete elements
type Book struct {
	Title string
	Price float64
}

func (b *Book) Accept(visitor Visitor) {
	visitor.VisitBook(b)
}

type Fruit struct {
	Name  string
	Price float64
	Qty   int
}

func (f *Fruit) Accept(visitor Visitor) {
	visitor.VisitFruit(f)
}

// Visitor interface
type Visitor interface {
	VisitBook(book *Book)
	VisitFruit(fruit *Fruit)
}

// Concrete visitor: Price calculator
type PriceCalculator struct {
	Total float64
}

func (p *PriceCalculator) VisitBook(book *Book) {
	p.Total += book.Price
}

func (p *PriceCalculator) VisitFruit(fruit *Fruit) {
	p.Total += fruit.Price * float64(fruit.Qty)
}

func main() {
	items := []Item{
		&Book{Title: "Go Design Patterns", Price: 50},
		&Fruit{Name: "Apple", Price: 2, Qty: 10},
	}

	calculator := &PriceCalculator{}

	for _, item := range items {
		item.Accept(calculator)
	}

	fmt.Println("Total price:", calculator.Total) // 50 + 2*10 = 70
}
