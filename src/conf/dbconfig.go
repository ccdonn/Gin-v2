package conf

import (
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
//	"log"
//	"reflect"
  . "utils"
)

var db *sql.DB

func InitDb() {
  var err error
  db, err = sql.Open("mysql", "root:45645655@tcp(127.0.0.1:3306)/zipdb")
  CheckErr(err, "Connection Fail")
  db.SetMaxIdleConns(0)
}


func GetConn() *sql.DB {
  return db
}