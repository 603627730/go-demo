package config

import (
	"github.com/gomodule/redigo/redis"
)

var RedisConn redis.Conn

func init() {
	c, err := redis.Dial("tcp", "localhost:6379",
		redis.DialPassword("lilishop"))

	if err != nil {
		// handle error
		panic(err)
	}
	RedisConn = c
	// defer c.Close()
}
