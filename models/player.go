package models

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