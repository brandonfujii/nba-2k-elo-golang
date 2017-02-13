package main 

import (
  "fmt"
  "../models"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
  _models := []interface{} {
    &models.Player{},
    &models.Game{}}

  for _, model := range _models {
    if !models.DB.HasTable(model) {
      fmt.Println("Creating %T table...", model)
      models.DB.CreateTable(model)
      fmt.Println("Created %T table", model)
    } else {
      fmt.Println("%T table already created", model)
    }
  }
}