package db

import (
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

type Repo struct {
	conn redis.Conn
}

func Connect() *Repo {
	pool := newPool()
	return &Repo{conn: pool.Get()}
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":"+os.Getenv("REDIS_PORT"))
			if err != nil {
				log.Panic(err)
			}

			return c, nil
		},
	}
}
