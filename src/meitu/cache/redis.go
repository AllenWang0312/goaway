package cache

import (
	"../../conf"
	model "../../model/meituri"
	"github.com/garyburd/redigo/redis"
	"time"
)

//func main() {

//}
var err error
var RedisClient *redis.Pool
var conn redis.Conn

func InitConn() {
	RedisClient = &redis.Pool{
		MaxIdle:     100,
		MaxActive:   4000,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			if conn, err = redis.Dial("tcp", conf.RediaHost+":6379"); err != nil {
				return nil, err
			}
			if _, err := conn.Do("AUTH", conf.RediaPass); err != nil {
				return nil, err
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}

	//defer conn.Close()
}

func SetV(k string, v string) error {
	rd:= RedisClient.Get()
	_, err := rd.Do("SET", k, v)
	if err != nil {
		return err
	}
	return nil
}
func SetBean(k string, v model.User) error {
	rd:= RedisClient.Get()
	_, err := rd.Do("SET", k, v)
	if err != nil {
		return err
	}
	return nil
}

func Set(k string, v string, sec uint64) error {
	rd:= RedisClient.Get()
	_, err := rd.Do("SET", k, v)
	if err != nil {
		return err
	}
	_, err = rd.Do("expire", k, sec) //10秒过期
	if err != nil {
		return err
	}
	return nil
}
func Get(k string) (string, error) {
	rd:= RedisClient.Get()
	v, err := redis.String(rd.Do("GET", k))
	return v, err
}

func GetSecondaryToken(k string) (string, error) {
	first, err := Get(k)
	if err == nil {
		println(first)
		return Get(first)
	} else {
		return "", err
	}
}
