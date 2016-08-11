package repository

import (
  "log"
//  "reflect"
  "database/sql"
  . "conf"
)

func FeedGet() *sql.Rows {
  db := GetConn()
  log.Println(db)
  rows, err := db.Query("select ID id, Title name from FEED")
  
//  log.Println(reflect.TypeOf(rows))
	
	if err != nil {
		log.Fatal(err)
		return nil
	}
	
	return rows
}