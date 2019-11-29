package cache

import (
	"../../conf"
	model "../model/meituri"
	"github.com/garyburd/redigo/redis"
)

//func main() {

//}
var conn redis.Conn

func InitConn() {
	conn, _ = redis.Dial("tcp", conf.RediaHost+":6379")
	//if err != nil {
	//	fmt.Println("connect redis error :", err)
	//	return
	//}
	//defer conn.Close()
}

func SetV(k string, v string) error {
	_, err := conn.Do("SET", k, v)
	if err != nil {
		return err
	}
	return nil
}
func SetBean(k string, v model.User) error {
	_, err := conn.Do("SET", k, v)
	if err != nil {
		return err
	}
	return nil
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
