package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse() // parses flags set above, default values will be used if not called

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// strip /static from incoming requests that match this route
	// so file searches within the subtree have the correct path
	// i.e. we don't want to find with a path starting with /static
	// whilst the current directory is /static
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Printf("starting server on %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
