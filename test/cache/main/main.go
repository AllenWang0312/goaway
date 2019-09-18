package main

import (
	//"../../cache"
	"fmt"
	"github.com/garyburd/redigo/redis"
)
var Pool redis.Pool
func init(){
	Pool = redis.Pool{
		MaxIdle:     16,
		MaxActive:   32,
		IdleTimeout: 120,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "10.1.210.69:6379")
		},
	}
}
func main() {
	//cache.Init()
	//err := cache.Set("token", "{josn}",1000)
	//if nil != err {
	//	panic(err)
	//}
	conn:=Pool.Get()
	res,err := conn.Do("HSET","student","name","jack")
	fmt.Println(res,err)
	res1,err := redis.String(conn.Do("HGET","student","name"))
	fmt.Printf("res:%s,error:%v",res1,err)
}
