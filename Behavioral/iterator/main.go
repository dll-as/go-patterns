package main

import "fmt"

// Item
type Book struct {
	Title string
}

// Iterator interface
type Iterator interface {
	HasNext() bool
	Next() *Book
}

// Aggregate interface
type BookCollection interface {
	CreateIterator() Iterator
}

// Concrete collection
type MyBooks struct {
	books []*Book
}

func (c *MyBooks) Add(book *Book) {
	c.books = append(c.books, book)
}

func (c *MyBooks) CreateIterator() Iterator {
	return &BookIterator{collection: c}
}

// Concrete iterator
type BookIterator struct {
	collection *MyBooks
	index      int
}

func (it *BookIterator) HasNext() bool {
	return it.index < len(it.collection.books)
}

func (it *BookIterator) Next() *Book {
	// if !it.HasNext() {
	// 	return nil
	// }

	book := it.collection.books[it.index]
	it.index++
	return book
}

func main() {
	collection := &MyBooks{}
	collection.Add(&Book{"Golang 101"})
	collection.Add(&Book{"Design Patterns"})
	collection.Add(&Book{"Clean Code"})

	iterator := collection.CreateIterator()

	for iterator.HasNext() {
		book := iterator.Next()
		fmt.Println("Book:", book.Title)
	}
}
