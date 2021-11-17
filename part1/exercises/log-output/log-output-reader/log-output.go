package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadFile(filepath.Join("/", "usr", "src", "app", "files", "logs.log"))
		if err != nil {
			fmt.Print(err)
		}

		logs := string(b)
		fmt.Fprintf(w, "<html>%s</html>", logs)
	})
	fmt.Println("Server started in 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}