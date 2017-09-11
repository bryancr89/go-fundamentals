package controller

import (
	"html/template"
	"net/http"
	"fmt"
	"webserver/model"
	"regexp"
	"strconv"
	"webserver/viewModels"
	"encoding/json"
	"log"
)

type books struct {
	booksTemplate *template.Template
}

func (b books) registerRoutes() {
	http.HandleFunc("/books", b.handleBooks)
	http.HandleFunc("/books/", b.handleBooks)
	http.HandleFunc("/api/books", b.handleNewBook)
}

func (b books) handleBooks(w http.ResponseWriter, request *http.Request) {
	bookPattern, _ := regexp.Compile(`/books/(\d+)`)
	matches := bookPattern.FindStringSubmatch(request.URL.Path)
	w.Header().Add("Content-type", "text/html")
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

func (b books) handleNewBook(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	book :=  model.Book{}
	err := dec.Decode(&book)
	if err != nil {
		log.Println(fmt.Errorf("Error retrieving books: %v", err))
		enc := json.NewEncoder(w)
		enc.Encode(viewModels.Book{})
		return
	}
	log.Println("Book:", book, r.Body)
	vm := viewModels.NewBooks([]model.Book{book})
	enc := json.NewEncoder(w)
	enc.Encode(vm)
}