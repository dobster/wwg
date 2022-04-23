package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	hostname = "localhost"
	port     = 8080
)

func main() {

	http.HandleFunc("/", handler)

	addr := fmt.Sprintf("%s:%d", hostname, port)
	log.Printf("Starting server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello user!")
	log.Println("The user called to say hello.")
}
