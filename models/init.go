package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "../config"
  "log"
)

type Model struct {
  ID uint `gorm:"primary_key; AUTO_INCREMENT" json:"id"`
}

var DB *gorm.DB 

func init() {
  var err interface{}
  DB, err = gorm.Open("mysql", config.Env.DBUser + ":" + config.Env.DBPassword + "@/" +  config.Env.DBName + "?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic(err) // AAAAAAAAA
  }
}

/* When passed a model type and ID, finds particular model with 
 * ID parameter and mutates the given model accordingly
 */ 
func GetByID(model interface{}, id uint) {
  DB.First(model, id)

  if DB.NewRecord(model) {
    log.Fatalln("Could not find model with that ID")
  }
}