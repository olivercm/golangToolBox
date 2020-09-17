package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//keys命令在match的数据量大的时候，会让redis卡顿， o(n)。scan使用iter，o(1)
func (db *DB) GetKeys(pattern string) ([]string, error) {
	iter := 0
	keys := make([]string, 0)
	for {
		arr, err := redis.Values(db.Do("SCAN", iter, "MATCH", pattern))
		if err != nil {
			return keys, fmt.Errorf("error retrieving '%s' keys", pattern)
		}

		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)

		if iter == 0 {
			break
		}
	}

	return keys, nil
}
