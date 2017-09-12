package with_go_redis

import (
	"github.com/go-redis/redis"
	"testing"
)

func BenchmarkWriteThroughSocket(b *testing.B) {
	client := redis.NewClient(&redis.Options{
		Network: "unix",
		Addr:    "/var/run/redis/redis.sock",
		DB:      1,
	})
	for i := 0; i < b.N; i++ {
		Write(client, "Socket", i)
	}
}

func BenchmarkWriteThroughTCP(b *testing.B) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   1,
	})
	for i := 0; i < b.N; i++ {
		Write(client, "TCP", i)
	}
}

func BenchmarkReadThroughSocket(b *testing.B) {
	client := redis.NewClient(&redis.Options{
		Network: "unix",
		Addr:    "/var/run/redis/redis.sock",
		DB:      1,
	})
	for i := 0; i < b.N; i++ {
		Read(client, "Socket", i)
	}
}

func BenchmarkReadThroughTCP(b *testing.B) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   1,
	})
	for i := 0; i < b.N; i++ {
		Read(client, "TCP", i)
	}
}
