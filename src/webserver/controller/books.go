package controller

import (
	"html/template"
	"net/http"
	"fmt"
	"webserver/model"
	"regexp"
	"strconv"
	"webserver/viewModels"
)

type books struct {
	booksTemplate *template.Template
}

func (b books) registerRoutes() {
	http.HandleFunc("/books", b.handleBooks)
	http.HandleFunc("/books/", b.handleBooks)
}

func (b books) handleBooks(w http.ResponseWriter, request *http.Request) {
	bookPattern, _ := regexp.Compile(`/books/(\d+)`)
	matches := bookPattern.FindStringSubmatch(request.URL.Path)
	if len(matches) > 0 {
		bookId, _ := strconv.Atoi(matches[1])
		b.handleBook(w, request, bookId)
	} else {
		books := model.GetBooks()
		vm := viewModels.NewBooks(books)
		b.booksTemplate.Execute(w, vm)
	}
}

func (b books) handleBook(w http.ResponseWriter, r *http.Request, bookId int) {
	book, err := model.GetBook(bookId)
	if err != nil {
		fmt.Println(err)
	}
	vm := viewModels.NewBooks([]model.Book{*book})
	b.booksTemplate.Execute(w, vm)
}