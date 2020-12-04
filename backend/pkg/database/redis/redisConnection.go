package redisConnection

import (
	"github.com/garyburd/redigo/redis"
)

var (
	PubSubConnection *redis.PubSubConn
	RedisConn        = func() (redis.Conn, error) {
		return redis.Dial("tcp", ":6364")
	}
)
