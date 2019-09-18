package cache

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var conn redis.Conn

func Init() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :", err)
		return
	}
	defer conn.Close()
}

func Set(k string, v string, sec uint64) error {
	_, err := conn.Do("SET", k, v)
	if err != nil {
		return err
	}
	_, err = conn.Do("expire", k, sec) //10秒过期
	if err != nil {
		return err
	}
	return nil
}
func Get(k string) (string, error) {
	v, err := redis.String(conn.Do("GET", k))
	if err != nil {
		return "", err
	} else {
		return v, nil
	}
}
