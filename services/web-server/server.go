package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server started in 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}