package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type DB struct {
	pool           redis.Pool
	scheme         string
	addr           string
	auth           string
	connectTimeout time.Duration
	readTimeout    time.Duration
	writeTimeout   time.Duration
}

func (db *DB) Do(cmd string, args ...interface{}) (interface{}, error) {
	conn := db.pool.Get()
	if conn.Err() != nil {
		return nil, conn.Err()
	}
	reply, err := conn.Do(cmd, args...)
	conn.Close()
	return reply, err
}

