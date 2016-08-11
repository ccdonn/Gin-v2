package utils

import (
  "log"
)

func CheckErr(err error, msg string) {
	if err != nil {
//		log.Panicf("%s", msg)
    log.Println("%s", msg)
	}
}