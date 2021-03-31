package di

import (
	"github.com/go-redis/redis/v8"
	"github.com/mix-go/xdi"
	"github.com/mix-go/dotenv"
	"time"
)

func init() {
	obj := xdi.Object{
		Name: "redis",
		New: func() (i interface{}, e error) {
			opt := redis.Options{
				Addr:        dotenv.Getenv("REDIS_ADDR").String(),
				Password:    dotenv.Getenv("REDIS_PASSWORD").String(),
				DB:          int(dotenv.Getenv("REDIS_DATABASE").Int64()),
				DialTimeout: time.Duration(dotenv.Getenv("REDIS_DIAL_TIMEOUT").Int64(10)) * time.Second,
			}
			return redis.NewClient(&opt), nil
		},
		Singleton: true,
	}
	if err := xdi.Provide(&obj); err != nil {
		panic(err)
	}
}

func Redis() (client *redis.Client) {
	if err := xdi.Populate("redis", &client); err != nil {
		panic(err)
	}
	return
}
