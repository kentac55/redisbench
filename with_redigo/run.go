package with_redigo

import "github.com/garyburd/redigo/redis"

func WriteDo(c redis.Conn, fix string, serial int) {
	_, err := c.Do("SET", "Redigo_keyWith"+fix+string(serial), serial)
	if err != nil {
		panic(err)
	}
}
func ReadDo(c redis.Conn, fix string, serial int) {
	_, err := c.Do("GET", "Redigo_keyWith"+fix+string(serial))
	if err != nil {
		panic(err)
	}
}
