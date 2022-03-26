package db

import (
	"os"

	"github.com/boltdb/bolt"
	log "github.com/sirupsen/logrus"
)

type Repo struct {
	DB *bolt.DB
}

const (
	bucket = "fs"
)

func Connect() *Repo {
	db, err := bolt.Open(os.Getenv("DB_FILE"), 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})

	if err != nil {
		log.Panic(err)
	}

	return &Repo{db}
}
