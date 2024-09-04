package main

import (
	"fmt"
	"html/template"
	"net/http"

	"rahuljuliato.com/go/museum/api"
	"rahuljuliato.com/go/museum/data"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PONG!"))
}

func handleTemplate(w http.ResponseWriter, R *http.Request) {
	html, err := template.ParseFiles("templates/index.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error!"))
		return
	}

	html.Execute(w, data.GetAll())
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/ping", handlePing)

	server.HandleFunc("/template", handleTemplate)

	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)

	// After go 1.22 you can register the method as such:
	// server.HandleFunc("GET /api/exhibitions", api.Get)
	// server.HandleFunc("POST /api/exhibitions/new", api.Post)

	fs := http.FileServer(http.Dir("./public/"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":3333", server)

	if err != nil {
		fmt.Println("Error while opening the server: ", err)
	}
}
