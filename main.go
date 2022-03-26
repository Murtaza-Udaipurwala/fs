package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/murtaza-udaipurwala/fs/api"
	"github.com/murtaza-udaipurwala/fs/db"
	"github.com/murtaza-udaipurwala/fs/shred"
	log "github.com/sirupsen/logrus"
)

func setup() error {
	godotenv.Load(".env")
	return os.MkdirAll("uploads", 0700)
}

func main() {
	if err := setup(); err != nil {
		log.Panic(err)
	}

	log.Infof("REDIS_PORT: %s\n", os.Getenv("REDIS_PORT"))
	log.Infof("PORT: %s\n", os.Getenv("PORT"))
	log.Infof("BASE_URL: %s\n", os.Getenv("BASE_URL"))

	dbR := db.Connect()
	dbS := db.NewService(dbR)
	apiS := api.NewService(*dbS)

	go shred.Start(apiS, dbS)

	api.Serve(apiS)
}
