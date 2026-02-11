package main

import "fmt"

// Small interfaces
type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type Faxer interface {
	Fax()
}

// Concrete types
type AllInOne struct{}

func (a *AllInOne) Print() { fmt.Println("AllInOne printing") }
func (a *AllInOne) Scan()  { fmt.Println("AllInOne scanning") }
func (a *AllInOne) Fax()   { fmt.Println("AllInOne faxing") }

type SimplePrinter struct{}

func (s *SimplePrinter) Print() { fmt.Println("SimplePrinter printing") }

// Client functions use only needed interfaces
func UsePrinter(p Printer) {
	p.Print()
}

func UseScanner(s Scanner) {
	s.Scan()
}

func main() {
	all := &AllInOne{}
	simple := &SimplePrinter{}

	UsePrinter(all)    // AllInOne printing
	UseScanner(all)    // AllInOne scanning
	UsePrinter(simple) // SimplePrinter printing
}
