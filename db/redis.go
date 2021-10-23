package db

import (
  "fmt"
  "github.com/go-redis/redis"
  "log"
)

var RedisClient *redis.Client

func InitRedisDb(host string, port string, password string, database string) {
    RedisClient = redis.NewClient(&redis.Options{
        Addr: fmt.Sprintf("%s:%s", host, port),
        Password: password,
        DB: 0,
    })
    if pong, err := RedisClient.Ping().Result(); err != nil {
        log.Printf("Failed to connect to Redis. %s", err)
        log.Panic("Could not connect to Redis")
    } else {
        fmt.Println("Redis connection successful. ", pong)
    }

}
