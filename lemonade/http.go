package lemonade

import "net/http"

func RunHttpServer() (err error) {
	server := &http.Server{
		Addr: ":8080",
		Handler: routers(),
	}

	return server.ListenAndServe()
}