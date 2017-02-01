package models

import (
  "os"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)


func init() {
  db, err := gorm.Open("mysql", os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@/2K?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic(err) // AAAAAAAAA
  }
}

