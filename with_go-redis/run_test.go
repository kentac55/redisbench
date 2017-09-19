package with_go_redis

import (
	"github.com/go-redis/redis"
	"testing"
)

var unixConf = &redis.Options{
	Network: "unix",
	Addr:    "/var/run/redis/redis.sock",
	DB:      1,
}

var tcpConf = &redis.Options{
	Addr: "localhost:6379",
	DB:   1,
}

func BenchmarkWriteSingleThroughSocket(b *testing.B) {
	client := redis.NewClient(unixConf)
	for i := 0; i < b.N; i++ {
		WriteSingle(client, "Socket", i)
	}
}

func BenchmarkWriteSingleThroughTCP(b *testing.B) {
	client := redis.NewClient(tcpConf)
	for i := 0; i < b.N; i++ {
		WriteSingle(client, "TCP", i)
	}
}

func BenchmarkReadSingleThroughSocket(b *testing.B) {
	client := redis.NewClient(unixConf)
	for i := 0; i < b.N; i++ {
		ReadSingle(client, "Socket", i)
	}
}

func BenchmarkReadSingleThroughTCP(b *testing.B) {
	client := redis.NewClient(tcpConf)
	for i := 0; i < b.N; i++ {
		ReadSingle(client, "TCP", i)
	}
}

func BenchmarkWritePipeWithSocket(b *testing.B) {
	client := redis.NewClient(unixConf)
	pipe := client.Pipeline()
	for i := 0; i < b.N; i++ {
		WritePipe(pipe, "Socket", i)
	}
	_, err := pipe.Exec()
	if err != nil {
		panic(err)
	}
}

func BenchmarkWritePipeWithTCP(b *testing.B) {
	client := redis.NewClient(tcpConf)
	pipe := client.Pipeline()
	for i := 0; i < b.N; i++ {
		WritePipe(pipe, "TCP", i)
	}
	_, err := pipe.Exec()
	if err != nil {
		panic(err)
	}
}

func BenchmarkReadPipeWithSocket(b *testing.B) {
	client := redis.NewClient(unixConf)
	pipe := client.Pipeline()
	for i := 0; i < b.N; i++ {
		ReadPipe(pipe, "Socket", i)
	}
	_, err := pipe.Exec()
	if err != nil {
		panic(err)
	}
}

func BenchmarkReadPipeWithTCP(b *testing.B) {
	client := redis.NewClient(tcpConf)
	pipe := client.Pipeline()
	for i := 0; i < b.N; i++ {
		ReadPipe(pipe, "TCP", i)
	}
	_, err := pipe.Exec()
	if err != nil {
		panic(err)
	}
}

func BenchmarkWriteTxWithSocket(b *testing.B) {
	client := redis.NewClient(unixConf)
	pipe := client.TxPipeline()
	for i := 0; i < b.N; i++ {
		WritePipe(pipe, "Socket", i)
	}
	_, err := pipe.Exec()
	if err != nil {
		panic(err)
	}
}

func BenchmarkWriteTxWithTCP(b *testing.B) {
	client := redis.NewClient(tcpConf)
	pipe := client.TxPipeline()
	for i := 0; i < b.N; i++ {
		WritePipe(pipe, "TCP", i)
	}
	_, err := pipe.Exec()
	if err != nil {
		panic(err)
	}
}

func BenchmarkReadTxWithSocket(b *testing.B) {
	client := redis.NewClient(unixConf)
	pipe := client.TxPipeline()
	for i := 0; i < b.N; i++ {
		ReadPipe(pipe, "Socket", i)
	}
	_, err := pipe.Exec()
	if err != nil {
		panic(err)
	}
}

func BenchmarkReadTxWithTCP(b *testing.B) {
	client := redis.NewClient(tcpConf)
	pipe := client.TxPipeline()
	for i := 0; i < b.N; i++ {
		ReadPipe(pipe, "TCP", i)
	}
	_, err := pipe.Exec()
	if err != nil {
		panic(err)
	}
}
