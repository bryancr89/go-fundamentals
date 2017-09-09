package model

import (
	"errors"
	"fmt"
)

type Book struct {
	Id int
	Title string
	Img string
	Categories[] Category
	Description string
	Published string
}

type Category struct {
	Name string
}

var elonMuskBook = Book{
	Id: 1,
	Title: "Elon Musk",
	Img: "https://www.manufacturing.net/sites/manufacturing.net/files/elon%20musk%20book%20jacket%202.jpg",
	Description: "Elon Musk: Tesla, SpaceX, and the Quest for a Fantastic Future is Ashlee Vance's biography of " +
	"Elon Musk, published in 2015...",
	Categories: []Category{
		Category{Name:"Biography"},
	},
	Published: "May 19, 2015",
}

var secretsOfJSNinja = Book{
	Id: 2,
	Title: "Secrets of the JavaScript Ninja",
	Img: "https://images-na.ssl-images-amazon.com/images/I/51UYwOhvQPL._SX258_BO1,204,203,200_.jpg",
	Description: "Summary Secrets of the Javascript Ninja takes you on a journey towards mastering modern " +
	"JavaScript development in three phases: design, construction, and maintenance...",
	Categories: []Category{
		Category{Name:"Javascript"},
		Category{Name:"Software Development"},
	},
	Published: "2008",
}

var books = []Book{
	elonMuskBook,
	secretsOfJSNinja,
}

func GetBooks() []Book {
	return books
}

func GetBook(id int) (*Book, error) {
	for _, book := range books {
		if book.Id == id {
			return &book, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Book with ID %v not found", id))
}