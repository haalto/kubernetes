package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1>", r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started in 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}