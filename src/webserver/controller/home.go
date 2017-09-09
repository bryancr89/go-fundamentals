package controller

import (
	"html/template"
	"net/http"
	"webserver/viewModels"
)

type home struct {
	homeTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, request *http.Request) {
	vm := viewModels.NewBase()
	h.homeTemplate.Execute(w, vm)
}