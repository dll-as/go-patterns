package main

import "fmt"

// Expression interface
type Expression interface {
	Interpret() int
}

// Number
type Number struct {
	value int
}

func (n *Number) Interpret() int {
	return n.value
}

// Add expression
type Add struct {
	left  Expression
	right Expression
}

func (a *Add) Interpret() int {
	return a.left.Interpret() + a.right.Interpret()
}

// Subtract expression
type Subtract struct {
	left  Expression
	right Expression
}

func (s *Subtract) Interpret() int {
	return s.left.Interpret() - s.right.Interpret()
}

func main() {
	// Represents the expression: (5 + 10) - 3
	expr := &Subtract{
		left: &Add{
			left:  &Number{5},
			right: &Number{10},
		},
		right: &Number{3},
	}

	result := expr.Interpret()
	fmt.Println("Result:", result) // 12
}
