package main

import (
	"net/http"
	"fmt"
	"os"
	"log"
	"io"
	"strings"
)

type myHandler struct {
	greeting string
}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", mh.greeting)))
}

func main() {
	http.ListenAndServe("8000", http.FileServer(http.Dir("public")))
}

func serveFilesUsingBuiltIn() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public" + r.URL.Path)
	})
	// Nil means DefaultServeMux
	http.ListenAndServe(":8000", nil)
}

func serveFiles() {
	http.Handle("/foo", &myHandler{greeting: "Hello"});

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("public" + r.URL.Path)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}

		defer f.Close()
		var contentType string
		switch {
		case strings.HasSuffix(r.URL.Path, "css"):
			contentType = "text/css"
		case strings.HasSuffix(r.URL.Path, "html"):
			contentType = "text/html"
		case strings.HasSuffix(r.URL.Path, "png"):
			contentType = "image/png"
		default:
			contentType = "text/plain"
		}
		w.Header().Add("Content-Type", contentType)
		io.Copy(w, f)

	})
}