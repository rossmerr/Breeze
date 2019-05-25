package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

var port int
var public string

func init() {
	flag.IntVar(&port, "port", 8080, "Port for http server")
	flag.StringVar(&public, "public", "../../public/", "Public folder for WASM")
}

func main() {
	flag.Parse()

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(""))))

	fs := http.FileServer(http.Dir(public))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
