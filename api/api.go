package api

import (
	"log"
	"net/http"
	"os"

	"github.com/murtaza-udaipurwala/fs/db"
)

func Serve(dbS db.Service) {
	s := NewService(dbS)
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
