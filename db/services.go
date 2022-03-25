package db

import "encoding/json"

func (db *DB) Set(key string, val interface{}) error {
	p, err := json.Marshal(val)
	if err != nil {
		return err
	}

	_, err = db.conn.Do("SET", key, p)
	return err
}
