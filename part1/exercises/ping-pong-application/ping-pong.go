package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var pongs = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pongs += 1
		fmt.Fprintf(w,"<h1>%s</h1>", "pong " + strconv.Itoa(pongs))
	})
	
	fmt.Println("Server started in 3001")
	log.Fatal(http.ListenAndServe(":3001", nil))
}