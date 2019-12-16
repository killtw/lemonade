package main

import (
	"github.com/killtw/lemonade/lemonade"
	"log"
)

func main() {
	if err := lemonade.RunHttpServer(); err != nil {
		log.Fatal(err)
	}
}

