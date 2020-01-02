package configs

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var (
	// DBRedis variable connection
	DBRedis *redis.Client
)

func init() {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: fmt.Sprintf("%s", password),
		DB:       0,
	})

	DBRedis = client
}
