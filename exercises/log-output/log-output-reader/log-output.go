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

		files, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }
 
    for _, f := range files {
            fmt.Println(f.Name())
    }

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