package db

import "errors"

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
