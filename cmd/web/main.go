package main

import (
	"flag"
	"log"
	"net/http"
)

type config struct {
	addr      string
	staticDir string
}

func main() {
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("starting server on %s", cfg.addr)

	err := http.ListenAndServe(cfg.addr, mux)
	log.Fatal(err)
}
