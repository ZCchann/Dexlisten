package db

import (
	"coin/apps/conf"
	"coin/pkg/log"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)
var RDB = new(redis.Client)

func ConnectRedis() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	RDB = redis.NewClient(&redis.Options{
		Addr:     conf.Conf().Redis.Addr,
		Password: conf.Conf().Redis.Password,
		DB:       conf.Conf().Redis.Database,
	})
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		log.Error(err)
		panic(err)
	}
}

