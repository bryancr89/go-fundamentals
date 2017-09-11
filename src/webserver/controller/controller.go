package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController home
	booksController books
)

func Startup(template *template.Template) {
	homeController.homeTemplate = template.Lookup("home.html")
	homeController.loginTemplate = template.Lookup("login.html")
	homeController.registerRoutes()
	booksController.booksTemplate = template.Lookup("books.html")
	booksController.registerRoutes()
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
}
