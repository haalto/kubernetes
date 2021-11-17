package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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

		f, err := os.OpenFile(filepath.Join("/", "usr", "src", "app", "files", "logs.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    if _, err := f.Write([]byte(t.Format(time.RFC850) + ": " + id.String() + "\n")); err != nil {
        log.Fatal(err)
    }
    if err := f.Close(); err != nil {
        log.Fatal(err)
    }

		fmt.Println(t.Format(time.RFC850) + ": " + id.String())
	}
}