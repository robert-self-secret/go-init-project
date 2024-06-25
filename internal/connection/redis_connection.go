package connection

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/robert-self-secret/go-init-project.git/internal/config"
)

func GetRedisConnection(cnf *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cnf.Redis.Address,
		Password: cnf.Redis.Password,
		DB:       0,
	})

	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)

	return rdb
}
