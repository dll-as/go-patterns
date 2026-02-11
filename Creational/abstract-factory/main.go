package main

import "fmt"

// Product interfaces
type DBConnection interface {
	Connect()
}

type DBQuery interface {
	Execute(query string)
}

// MySQL implementations
type MySQLConnection struct{}

func (m *MySQLConnection) Connect() {
	fmt.Println("Connected to MySQL")
}

type MySQLQuery struct{}

func (m *MySQLQuery) Execute(query string) {
	fmt.Println("MySQL executing:", query)
}

// PostgreSQL implementations
type PostgresConnection struct{}

func (p *PostgresConnection) Connect() {
	fmt.Println("Connected to PostgreSQL")
}

type PostgresQuery struct{}

func (p *PostgresQuery) Execute(query string) {
	fmt.Println("Postgres executing:", query)
}

// Abstract Factory interface
type DatabaseFactory interface {
	CreateConnection() DBConnection
	CreateQuery() DBQuery
}

// Concrete factories
type MySQLFactory struct{}

func (m *MySQLFactory) CreateConnection() DBConnection {
	return &MySQLConnection{}
}

func (m *MySQLFactory) CreateQuery() DBQuery {
	return &MySQLQuery{}
}

type PostgresFactory struct{}

func (p *PostgresFactory) CreateConnection() DBConnection {
	return &PostgresConnection{}
}

func (p *PostgresFactory) CreateQuery() DBQuery {
	return &PostgresQuery{}
}

func main() {
	var factory DatabaseFactory

	dbType := "mysql" // change to "postgres"

	if dbType == "mysql" {
		factory = &MySQLFactory{}
	} else {
		factory = &PostgresFactory{}
	}

	conn := factory.CreateConnection()
	query := factory.CreateQuery()

	conn.Connect()
	query.Execute("SELECT * FROM users")
}
