package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/murtaza-udaipurwala/fs/api"
	"github.com/murtaza-udaipurwala/fs/db"
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

	r := db.Connect()
	s := db.NewService(r)
	api.Serve(*s)
}
