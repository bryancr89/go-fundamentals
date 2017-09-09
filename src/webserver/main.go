package main

import (
	"html/template"
	"net/http"
	"webserver/controller"
)

func main() {
	templates := populateTemplates()
	controller.Startup(templates)

	http.ListenAndServe(":8000", nil)
}

func populateTemplates() *template.Template {
	result := template.New("templates")
	const basePath = "templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}
