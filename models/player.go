package models

import "log"

type Player struct {
  Model
  Username string `gorm:"size:255; not null; unique" json:"username"`
  Rating int `json: "rating"`
}

func (p Player) ToJSON() interface{} {
  data, player := make(map[string]interface{}), make(map[string]interface{})

  player["id"] = p.ID
  player["username"] = p.Username
  player["rating"] = p.Rating

  data["player"] = player

  return data
}

func (p Player) GetById(id uint) Player {
  var player Player
  DB.First(&player, id)

  if DB.NewRecord(player) {
    log.Fatalln("Could not find player with that ID")
  }

  return player
}