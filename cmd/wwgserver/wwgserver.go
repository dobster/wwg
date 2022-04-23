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

	http.HandleFunc("/", HelloWorldHandler)

	addr := fmt.Sprintf("%s:%d", hostname, port)
	log.Printf("Starting server on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
	log.Println("The user called to say hello.")
}
