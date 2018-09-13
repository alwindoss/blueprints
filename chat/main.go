package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
	"time"
)

type templateHandler struct {
	once     sync.Once
	fileName string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.fileName)))
	})
	d := struct {
		Title   string
		Name    string
		TimeNow string
	}{
		"Chat App",
		"Alwin",
		time.Now().String(),
	}
	t.templ.Execute(w, d)
}

func main() {
	log.Println("Starting Chat Server")
	http.Handle("/", &templateHandler{fileName: "chat.html"})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
