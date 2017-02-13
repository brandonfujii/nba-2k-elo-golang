package main 

import (
  "fmt"
  "reflect"

  "../models"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
  _models := []interface{} {
    &models.Player{},
    &models.Game{}}

  for _, model := range _models {
    if models.DB.HasTable(model) {
      models.DB.DropTable(model)
      fmt.Println("Dropped table for", reflect.TypeOf(model))
    }
  }

  for _, model := range _models {
    if !models.DB.HasTable(model) {
      fmt.Println("Creating table for ", reflect.TypeOf(model), "...")
      models.DB.CreateTable(model)
      fmt.Println("Created table", reflect.TypeOf(model))
    }
  }
}