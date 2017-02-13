package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "../config"
)

type Model struct {
  ID uint `gorm:"primary_key; AUTO_INCREMENT" json:"id"`
}

var DB *gorm.DB 

func init() {
  var err interface{}
  DB, err = gorm.Open("mysql", config.Env.DBUser + ":" + config.Env.DBPassword + "@/2K?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic(err) // AAAAAAAAA
  }
}