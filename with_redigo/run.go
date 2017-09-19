package with_redigo

import (
	"github.com/garyburd/redigo/redis"
)

func WriteSingle(c redis.Conn, fix string, serial int) {
	_, err := c.Do("SET", "Redigo_keyWith"+fix+string(serial), serial)
	if err != nil {
		panic(err)
	}
}

func ReadSingle(c redis.Conn, fix string, serial int) {
	_, err := c.Do("GET", "Redigo_keyWith"+fix+string(serial))
	if err != nil {
		panic(err)
	}
}

func WritePipeline(c redis.Conn, fix string, serial int) {
	c.Send("SET", "RedigoP_keyWith"+fix+string(serial), serial)
}

func ReadPipeline(c redis.Conn, fix string, serial int) {
	c.Send("GET", "RedigoP_keyWith"+fix+string(serial))
}

func WriteTx(c redis.Conn, fix string, serial int) {
	c.Send("SET", "RedigoTx_keyWith"+fix+string(serial), serial)
}

func ReadTx(c redis.Conn, fix string, serial int) {
	c.Send("GET", "RedigoTx_keyWith"+fix+string(serial))
}
