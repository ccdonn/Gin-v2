package main

import (
  "github.com/gin-gonic/gin"
  . "conf"
)

var router *gin.Engine

func init() {
  router = gin.Default()
  InitDb()
  InitCache()
  Routes()
}

func main() {
  router.Run()
}