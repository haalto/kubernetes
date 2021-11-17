package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pongs := 0

		b, err := ioutil.ReadFile(filepath.Join("/", "usr", "src", "app", "files", "pongs.txt"))
		if err != nil {
			log.Print(err)
		} else {
			p, err := strconv.Atoi(string(b))

			if (err != nil) {
				log.Print(err)
				p = 0
			} else {
				pongs = p
			}
		}

		if (err != nil) {
			log.Print(err)
		}

		pongs += 1

		err = ioutil.WriteFile(filepath.Join("/", "usr", "src", "app", "files", "pongs.txt"), []byte(strconv.Itoa(pongs)), 0644 )

		if (err != nil) {
			log.Print(err)
		}

		fmt.Fprintf(w,"<h1>%s</h1>", "pong " + strconv.Itoa(pongs))
	})
	
	fmt.Println("Server started in 3001")
	log.Fatal(http.ListenAndServe(":3001", nil))
}