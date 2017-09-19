package with_redigo

import (
	"github.com/garyburd/redigo/redis"
	"net"
	"testing"
	"time"
)

const UNIXADDR = "/var/run/redis/redis.sock"
const TIMEOUT = 10000000000

var unixConn, _ = net.Dial("unix", UNIXADDR)
var unixCli = redis.NewConn(unixConn, TIMEOUT, TIMEOUT)
var tcpCli, _ = redis.Dial("tcp", ":6379")
var unixPool = &redis.Pool{
	MaxIdle:     3,
	IdleTimeout: 240 * time.Second,
	Dial:        func() (redis.Conn, error) { return unixCli, nil },
}
var tcpPool = &redis.Pool{
	MaxIdle:     3,
	IdleTimeout: 240 * time.Second,
	Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", ":6379") },
}
var unixCliP = unixPool.Get()
var tcpCliP = unixPool.Get()

func BenchmarkWriteSingleThroughSocket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteSingle(unixCli, "Socket", i)
	}
}

func BenchmarkWriteSingleThroughTCP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteSingle(tcpCli, "TCP", i)
	}
}

func BenchmarkReadSingleThroughSocket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadSingle(unixCli, "Socket", i)
	}
}

func BenchmarkReadSingleThroughTCP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadSingle(tcpCli, "TCP", i)
	}
}

func BenchmarkWritePipelineThroughSocket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WritePipeline(unixCli, "Socket", i)
	}
	unixCli.Flush()
	_, err := unixCli.Receive()
	if err != nil {
		panic(err)
	}
}

func BenchmarkWritePipelineThroughTCP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WritePipeline(tcpCli, "TCP", i)
	}
	tcpCli.Flush()
	_, err := tcpCli.Receive()
	if err != nil {
		panic(err)
	}
}

func BenchmarkReadPipelineThroughSocket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadPipeline(unixCli, "Socket", i)
	}
	unixCli.Flush()
	_, err := unixCli.Receive()
	if err != nil {
		panic(err)
	}
}

func BenchmarkReadPipelineThroughTCP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadPipeline(tcpCli, "TCP", i)
	}
	tcpCli.Flush()
	_, err := tcpCli.Receive()
	if err != nil {
		panic(err)
	}
}

func BenchmarkWriteTxThroughSocket(b *testing.B) {
	unixCli.Send("MULTI")
	for i := 0; i < b.N; i++ {
		WriteTx(unixCli, "Socket", i)
	}
	_, err := unixCli.Do("EXEC")
	if err != nil {
		panic(err)
	}
}

func BenchmarkWriteTxThroughTCP(b *testing.B) {
	tcpCli.Send("MULTI")
	for i := 0; i < b.N; i++ {
		WriteTx(tcpCli, "TCP", i)
	}
	_, err := tcpCli.Do("EXEC")
	if err != nil {
		panic(err)
	}
}

func BenchmarkReadTxThroughSocket(b *testing.B) {
	unixCli.Send("MULTI")
	for i := 0; i < b.N; i++ {
		WriteTx(unixCli, "Socket", i)
	}
	_, err := unixCli.Do("EXEC")
	if err != nil {
		panic(err)
	}
}

func BenchmarkReadTxThroughTCP(b *testing.B) {
	tcpCli.Send("MULTI")
	for i := 0; i < b.N; i++ {
		WriteTx(tcpCli, "TCP", i)
	}
	_, err := tcpCli.Do("EXEC")
	if err != nil {
		panic(err)
	}
}

func BenchmarkWriteSingleThroughSocketPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteSingle(unixCliP, "Socket", i)
	}
}

func BenchmarkWriteSingleThroughTCPPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WriteSingle(tcpCliP, "TCP", i)
	}
}

func BenchmarkReadSingleThroughSocketPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadSingle(unixCliP, "Socket", i)
	}
}

func BenchmarkReadSingleThroughTCPPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadSingle(tcpCliP, "TCP", i)
	}
}

func BenchmarkWritePipelineThroughSocketPool(b *testing.B) {
	c := unixPool.Get()
	defer c.Close()
	for i := 0; i < b.N; i++ {
		WritePipeline(unixCliP, "Socket", i)
	}
	unixCliP.Flush()
	_, err := unixCliP.Receive()
	if err != nil {
		panic(err)
	}
}

func BenchmarkWritePipelineThroughTCPPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WritePipeline(tcpCliP, "TCP", i)
	}
	tcpCliP.Flush()
	_, err := tcpCliP.Receive()
	if err != nil {
		panic(err)
	}
}

func BenchmarkReadPipelineThroughSocketPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadPipeline(unixCliP, "Socket", i)
	}
	unixCliP.Flush()
	_, err := unixCliP.Receive()
	if err != nil {
		panic(err)
	}
}

func BenchmarkReadPipelineThroughTCPPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadPipeline(tcpCliP, "TCP", i)
	}
	tcpCliP.Flush()
	_, err := tcpCliP.Receive()
	if err != nil {
		panic(err)
	}
}
