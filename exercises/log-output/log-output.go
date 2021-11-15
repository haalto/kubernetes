package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	for {
		time.Sleep(5 * time.Second)
		fmt.Println(id.String())
	}
}