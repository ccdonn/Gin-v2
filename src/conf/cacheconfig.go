package conf

import (
  "github.com/garyburd/redigo/redis"
//	"log"
	. "utils"
)

var c redis.Conn

func InitCache() {
	var err error
	c, err = redis.Dial("tcp", "localhost:6379")
	CheckErr(err, "redis connection fail")
//  log.Println(reflect.TypeOf(c))
  _, err = c.Do("AUTH", "taiwan#1")
  CheckErr(err, "redis Auth Fail")
  
}


func GetCacheConn() redis.Conn {
  return c
}