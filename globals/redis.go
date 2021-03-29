package globals

import (
	"github.com/go-redis/redis/v8"
	"github.com/mix-go/xdi"
)

func Redis() (client *redis.Client) {
	if err := xdi.Populate("redis", &client); err != nil {
		panic(err)
	}
	return
}
