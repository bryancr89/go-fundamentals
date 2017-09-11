package main

import (
	"html/template"
	"net/http"
	"webserver/controller"
	"webserver/middleware"
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"fmt"
	"webserver/model"
)

func main() {
	templates := populateTemplates()
	db := connectToDatabase()
	defer db.Close()
	controller.Startup(templates)
	http.ListenAndServeTLS(":8000", "cert.pem", "key.pem", &middleware.TimeoutMiddleware{
		new(middleware.GzipMiddleware),
	})
}

func connectToDatabase() *sql.DB {
	db, err := sql.Open("postgres", "postgres://gotest:gotest@localhost/gotest?sslmode=disable")
	if err != nil {
		log.Fatalln(fmt.Errorf("unable to connect to database: %v", err))
	}
	model.SetDatabase(db)
	return db
}

func populateTemplates() *template.Template {
	result := template.New("templates")
	const basePath = "templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}
