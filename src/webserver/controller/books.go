package controller

import (
	"html/template"
	"net/http"
	"webserver/viewModels"
	"fmt"
)

type books struct {
	booksTemplate *template.Template
}

func (b books) registerRoutes() {
	http.HandleFunc("/books", b.handleBooks)
}

func (b books) handleBooks(w http.ResponseWriter, request *http.Request) {
	vm := viewModels.NewBooks()
	fmt.Println(vm)
	b.booksTemplate.Execute(w, vm)
}