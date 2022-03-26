package db

import (
	"errors"

	"github.com/boltdb/bolt"
)

var ErrDoesNotExist = errors.New("key does not exist")

func (r *Repo) Set(key string, val []byte) error {
	err := r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Put([]byte(key), val)
	})

	return err
}

func (r *Repo) Get(key string) ([]byte, error) {
	var val []byte

	err := r.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		val = b.Get([]byte(key))
		if val == nil {
			return ErrDoesNotExist
		}

		return nil
	})

	return val, err
}

func (r *Repo) Del(key string) error {
	err := r.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		return b.Delete([]byte(key))
	})

	return err
}

func (r *Repo) GetAll() ([]string, error) {
	var key []string

	err := r.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.ForEach(func(k, v []byte) error {
			key = append(key, string(k))
			return nil
		})

		return err
	})

	return key, err
}
