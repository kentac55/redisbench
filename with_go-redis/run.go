package with_go_redis

import (
	"github.com/go-redis/redis"
)

func Write(client *redis.Client, fix string, serial int) {
	err := client.Set("keyWith"+fix+string(serial), serial, 0).Err()
	if err != nil {
		panic(err)
	}
}

func Read(client *redis.Client, fix string, serial int) {
	_, err := client.Get("keyWithSock" + string(serial)).Result()
	if err != nil {
		panic(err)
	}
}
