package main

import (
	"github.com/killtw/lemonade/http"
	"github.com/killtw/lemonade/lemonade"
	"golang.org/x/sync/errgroup"
	"log"
)

func main() {
	var err error
	var e errgroup.Group

	e.Go(lemonade.InitTrie)
	e.Go(http.RunHttpServer)

	if err = e.Wait(); err != nil {
		log.Fatalln(err)
	}
}
