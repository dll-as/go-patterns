package main

import "fmt"

// Subject interface
type Database interface {
	Query(sql string)
}

// RealSubject
type RealDatabase struct{}

func (r *RealDatabase) Query(sql string) {
	fmt.Println("Executing query:", sql)
}

// Proxy
type DatabaseProxy struct {
	realDatabase *RealDatabase
	user         string
}

func NewDatabaseProxy(user string) *DatabaseProxy {
	return &DatabaseProxy{user: user}
}

func (p *DatabaseProxy) Query(sql string) {
	if !p.hasAccess() {
		fmt.Println("Access denied for user:", p.user)
		return
	}

	// Lazy initialization
	if p.realDatabase == nil {
		p.realDatabase = &RealDatabase{}
	}

	// Logging
	fmt.Println("User", p.user, "executing query")
	p.realDatabase.Query(sql)
}

func (p *DatabaseProxy) hasAccess() bool {
	// simple access check
	return p.user == "admin"
}

func main() {
	adminProxy := NewDatabaseProxy("admin")
	adminProxy.Query("SELECT * FROM users")

	guestProxy := NewDatabaseProxy("guest")
	guestProxy.Query("SELECT * FROM users")
}
