package cacheRepo

import (
  "github.com/garyburd/redigo/redis"
  . "conf"
  . "utils"
//  "log"
)

func Get(name string) string {
  conn := GetCacheConn()
  info, err := redis.String(conn.Do("GET", name))
	if err != nil {
    CheckErr(err, "redis get fail")
		return ""
	} else {
	  return info
	}
}