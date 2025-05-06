package initialize

import (
	"fmt"
	"github.com/go-redis/redis"
	"my-take-out/global"
)

func initRedis() *redis.Client {
	redisOpt := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisOpt.Host, redisOpt.Port),
		Password: redisOpt.Password,
		DB:       redisOpt.DataBase,
	})
	ping := client.Ping()
	err := ping.Err()
	if err != nil {
		panic(err)
	}
	return client
}
