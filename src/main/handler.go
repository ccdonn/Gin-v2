package main

import (
	"container/list"
	//	"fmt"
	cacheRepo "cacheRepo"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	repo "repository"
	. "response"
	"time"
)

func GetHelloV2(c *gin.Context) {
	c.String(200, "working")
}

func GetFeedWithIdv1(c *gin.Context) {

	var (
		rtext string
	)

	pathName := c.Param("name")
	log.Println("pathName: ", pathName)
	if len(pathName) > 0 {
		rtext = cacheRepo.Get(pathName)
	}

	if len(pathName) > 0 && len(rtext) > 0 {
		log.Println("pathName:", rtext)
		bytes := []byte(rtext)
		var redisres RedisRes
		json.Unmarshal(bytes, &redisres)

		c.JSON(200, FeedRedisResponse{Response: Response{Status: "success", Timestamp: time.Now().UTC().Unix()}, Result: redisres})
	} else {
		c.JSON(http.StatusBadRequest, Response{Status: "failure", Timestamp: time.Now().UTC().Unix()})
	}

}

func GetFeedv1(c *gin.Context) {

	var (
		id    string
		name  string
	)

	rows := repo.FeedGet()
	l := list.New()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Panic(err)
		}
		ob := Obje{
			Id:   id,
			Name: name,
		}

		l.PushBack(ob)
	}

	var slice []Obje = make([]Obje, l.Len())

	cnt := 0
	for iter := l.Front(); iter != nil; iter = iter.Next() {
		slice[cnt] = iter.Value.(Obje)
		cnt++
	}

	feedResponse := FeedResponse{
		Status:    "success",
		Timestamp: time.Now().Unix(),
		Result:    slice,
	}
	c.JSON(http.StatusOK, feedResponse)

}

func Getv2(c *gin.Context) {
	c.JSON(http.StatusOK, "getv2")
}

func GetLoopv2(c *gin.Context) {
  log.Println(c.Request.Method)
  log.Println(c.Request.Host)
  log.Println(c.Request.Header)
  
  rows := repo.FeedGet()
  log.Println(rows)
  for ;; {
    
  }
  
}

func undefineFunc(c *gin.Context) {
  c.String(200, "Undefined")
}