package db

import (
	"errors"

	"github.com/gomodule/redigo/redis"
)

var ErrDoesNotExist = errors.New("key does not exist")

func (r *Repo) Set(key string, val []byte) error {
	_, err := r.conn.Do("SET", key, val)
	return err
}

func (r *Repo) Get(key string) ([]byte, error) {
	val, err := r.conn.Do("GET", key)
	if err != nil {
		return nil, err
	}

	if val == nil {
		return nil, ErrDoesNotExist
	}

	return val.([]byte), nil
}

func (r *Repo) Del(key string) error {
	_, err := r.conn.Do("DEL", key)
	return err
}

func (r *Repo) Exists(key string) (bool, error) {
	return redis.Bool(r.conn.Do("EXISTS", key))
}

func (r *Repo) GetAll() ([]string, error) {
	return redis.Strings(r.conn.Do("KEYS", "*"))
}
