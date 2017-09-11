package main

import (
	"html/template"
	"net/http"
	"webserver/controller"
	"webserver/middleware"
)

func main() {
	templates := populateTemplates()
	controller.Startup(templates)
	http.ListenAndServe(":8000", new(middleware.GzipMiddleware))
}

func populateTemplates() *template.Template {
	result := template.New("templates")
	const basePath = "templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}
