package http

import "net/http"

func RunHttpServer() error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: routers(),
	}

	return server.ListenAndServe()
}
