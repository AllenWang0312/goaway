package cache
//
//import (
//	"../../conf"
//	"fmt"
//	"github.com/go-redis/redis"
//	"time"
//)
//
//var client *redis.Client
//
//func InitConn() {
//	client = redis.NewClient(&redis.Options{
//		Addr:     conf.RediaHost + ":6379",
//		Password: "", // no password set
//		DB:       0,          // use default DB
//	})
//	pong, err := client.Ping().Result()
//	fmt.Println(pong, err)
//}
//
//func Set(key string, value string, exp time.Duration) error {
//	err := client.Set(key, value, exp).Err()
//	return err
//}
//
//func Get(key string) string {
//	v, _ := client.Get(key).Result()
//	if len(v) > 0 {
//		return v
//	} else {
//		return ""
//	}
//}
