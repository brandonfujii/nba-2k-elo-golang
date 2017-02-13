package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
  ID uint `gorm:"primary_key; AUTO_INCREMENT" json:"id"`
}

const initialRating int = 1300
const kFactor float64 = 20.0
var DB *gorm.DB 

func init() {
  var err interface{}
  DB, err = gorm.Open("mysql", "root:danfibjr360!@/2K?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic(err) // AAAAAAAAA
  }
}