package api

import (
	"log"
	"net/http"
	"os"
)

func Serve(s *Service) {
	c := NewController(s)

	mux := http.NewServeMux()
	mux.HandleFunc("/", c.Handler)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	log.Printf("Listening on port :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
