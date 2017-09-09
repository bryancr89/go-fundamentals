package controller

import (
	"html/template"
	"net/http"
	"webserver/viewModels"
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
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewModels.NewLogin()
	h.loginTemplate.Execute(w, vm)
}