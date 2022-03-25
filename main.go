package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/murtaza-udaipurwala/fs/api"
	"github.com/murtaza-udaipurwala/fs/db"
	"github.com/murtaza-udaipurwala/fs/expire"
)

func setup() error {
	godotenv.Load(".env")
	return os.MkdirAll("uploads", 0700)
}

func main() {
	if err := setup(); err != nil {
		log.Panic(err)
	}

	log.Printf("REDIS_PORT: %s\n", os.Getenv("REDIS_PORT"))
	log.Printf("PORT: %s\n", os.Getenv("PORT"))
	log.Printf("BASE_URL: %s\n", os.Getenv("BASE_URL"))

	dbR := db.Connect()
	dbS := db.NewService(dbR)
	apiS := api.NewService(*dbS)

	go expire.Shred(apiS, dbS)

	api.Serve(apiS)
}
