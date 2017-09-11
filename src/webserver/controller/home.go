package controller

import (
	"html/template"
	"net/http"
	"webserver/viewModels"
	"fmt"
	"log"
)

type home struct {
	homeTemplate *template.Template
	loginTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/login", h.handleLogin)
}

func (h home) handleHome(w http.ResponseWriter, request *http.Request) {
	vm := viewModels.NewBase()
	w.Header().Add("Content-type", "text/html")
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewModels.NewLogin()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("Error logging in: %v", err))
		}
		email := r.Form.Get("email")
		password:= r.Form.Get("password")
		if email == "foo@example.com" {
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		}
		vm.Email = email
		vm.Password = password
	}
	w.Header().Add("Content-type", "text/html")
	h.loginTemplate.Execute(w, vm)
}