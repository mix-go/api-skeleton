package globals

import (
    "github.com/go-redis/redis/v8"
    "github.com/mix-go/console"
)

func Redis() *redis.Client {
    return console.Context().Get("redis").(*redis.Client)
}
