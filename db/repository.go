package db

func (r *Repo) Set(key string, val []byte) error {
	_, err := r.conn.Do("SET", key, val)
	return err
}

func (r *Repo) Get(key string) ([]byte, error) {
	val, err := r.conn.Do("GET", key)
	if err != nil {
		return nil, err
	}

	return val.([]byte), nil
}
