package main

import (
	"github.com/killtw/lemonade/http"
	"github.com/killtw/lemonade/lemonade"
	"github.com/killtw/lemonade/rpc"
	"golang.org/x/sync/errgroup"
	"log"
)

func main() {
	var err error
	var e errgroup.Group

	if err:= lemonade.InitTrie(); err != nil {
		log.Fatalln(err)
	}

	e.Go(http.RunHttpServer)
	e.Go(rpc.RunGRPCServer)

	if err = e.Wait(); err != nil {
		log.Fatalln(err)
	}
}
