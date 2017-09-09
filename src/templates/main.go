package main

// "html/template" & "text/template" are similar, the html one has logic that escape strings
// (for prevent security issues)
import (
	"html/template"
	"fmt"
	"os"
	"net/http"
	"log"
)

func main() {
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		requestedFile := request.URL.Path[1:]
		t := templates.Lookup(requestedFile + ".html")
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.ListenAndServe(":8000", nil)
}

func populateTemplates() *template.Template {
	result := template.New("templates")
	const basePath = "templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}

func basicTemplate() {
	templateString := `Lemonade STand Supply`
	t, err := template.New("title").Parse(templateString)

	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println(err)
	}
}
