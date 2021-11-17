package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/google/uuid"
)


func main() {
	generateLog()
}

func generateLog() {
	for {
		time.Sleep(5 * time.Second)
		id := uuid.New()
		t := time.Now()
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

		f, err := os.OpenFile(filepath.Join("/", "usr", "src", "app", "files", "logs.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    if _, err := f.Write([]byte(t.Format(time.RFC850) + ": " + id.String() + "\n" + "Ping / Pongs: " + strconv.Itoa(pongs) + "\n")); err != nil {
        log.Fatal(err)
    }
    if err := f.Close(); err != nil {
        log.Fatal(err)
    }

		fmt.Println(t.Format(time.RFC850) + ": " + id.String())
	}
}