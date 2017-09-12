package with_redigo

import (
	"github.com/garyburd/redigo/redis"
	"net"
	"testing"
)

func BenchmarkWriteThroughSocketDo(b *testing.B) {
	unixConn, err := net.Dial("unix", "/var/run/redis/redis.sock")
	if err != nil {
		panic(err)
	}
	c := redis.NewConn(unixConn, 100000000, 100000000)
	for i := 0; i < b.N; i++ {
		WriteDo(c, "Socket", i)
	}
}
func BenchmarkWriteThroughTCPDo(b *testing.B) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		WriteDo(c, "TCP", i)
	}
}
func BenchmarkReadThroughSocketDo(b *testing.B) {
	unixConn, err := net.Dial("unix", "/var/run/redis/redis.sock")
	if err != nil {
		panic(err)
	}
	c := redis.NewConn(unixConn, 100000000, 100000000)
	for i := 0; i < b.N; i++ {
		ReadDo(c, "Socket", i)
	}
}
func BenchmarkReadThroughTCPDo(b *testing.B) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		ReadDo(c, "TCP", i)
	}
}
