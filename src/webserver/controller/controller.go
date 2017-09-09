package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController home
)

func Startup(template *template.Template) {
	t := template.Lookup("home.html")
	homeController.homeTemplate = t
	homeController.registerRoutes()
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
