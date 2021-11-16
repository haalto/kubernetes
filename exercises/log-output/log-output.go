package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var id = uuid.New()
var t = time.Now()

func main() {
	id = uuid.New()
	t = time.Now()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>%s</h1>", t.Format(time.RFC850) + ": " + id.String())
	})
	go generateLog()
	fmt.Println("Server started in 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func generateLog() {
	for {
		time.Sleep(5 * time.Second)
		id = uuid.New()
		t = time.Now()
		fmt.Println(t.Format(time.RFC850) + ": " + id.String())
	}
}