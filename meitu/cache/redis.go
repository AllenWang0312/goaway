package cache
//
//import (
//	"../../conf"
//	"github.com/garyburd/redigo/redis"
//)
//
////func main() {
//
////}
//var conn redis.Conn
//
//func InitConn() {
//	conn, _ = redis.Dial("tcp", conf.RediaHost+":6379")
//	//if err != nil {
//	//	fmt.Println("connect redis error :", err)
//	//	return
//	//}
//	//defer conn.Close()
//}
//
//func Set(k string, v string, sec uint64) error {
//	_, err := conn.Do("SET", k, v)
//	if err != nil {
//		return err
//	}
//	_, err = conn.Do("expire", k, sec) //10秒过期
//	if err != nil {
//		return err
//	}
//	return nil
//}
//func Get(k string) (string, error) {
//	v, err := redis.String(conn.Do("GET", k))
//	if err != nil {
//		return "", err
//	} else {
//		return v, nil
//	}
//}
