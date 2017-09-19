package with_go_redis

import (
	"github.com/go-redis/redis"
)

func WriteSingle(client *redis.Client, fix string, serial int) {
	err := client.Set("keyWith"+fix+string(serial), serial, 0).Err()
	if err != nil {
		panic(err)
	}
}

func ReadSingle(client *redis.Client, fix string, serial int) {
	_, err := client.Get("keyWith" + fix + string(serial)).Result()
	if err != nil {
		panic(err)
	}
}

func WritePipe(pipe redis.Pipeliner, fix string, serial int) {
	pipe.Set("keyWith"+fix+string(serial), serial, 0)
}

func ReadPipe(pipe redis.Pipeliner, fix string, serial int) {
	pipe.Get("keyWith" + fix + string(serial))
}
