package redisConnection

import (
	"github.com/garyburd/redigo/redis"
)

var PubSubConnection *redis.PubSubConn

func RedisConn() (redis.Conn, error) {
	return redis.Dial("tcp", "redis:6379")
}
