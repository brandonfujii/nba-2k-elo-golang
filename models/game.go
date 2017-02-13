package models

import (
  "math"
  "time"
  "log"
)

const K_FACTOR float64 = 20.0

type Game struct {
  Model
  Winner Player `gorm:"ForeignKey:WinnerId" json:"winner"`
  Loser Player `gorm:"ForeignKey:LoserId" json:"loser"`

  WinnerId uint `json:"winner_id"`
  LoserId uint `json:"loser_id"`

  WinnerRating int `json:"loser_rating"`
  LoserRating int `json:"winner_rating"`

  WinnerPoints int `json:"winner_points"`
  LoserPoints int `json:"loser_points"`

  Date int `json:"date"`
}

/* Convert struct to JSON */
func (g *Game) ToJSON() interface{} {
  data, game := make(map[string]interface{}), make(map[string]interface{})

  game["id"] = g.ID
  game["winner_id"] = g.WinnerId
  game["loser_id"] = g.LoserId
  game["winner_rating"] = g.WinnerRating
  game["loser_rating"] = g.LoserRating
  game["date"] = g.Date

  data["game"] = game

  return data
}

func (g Game) GetById(id uint) Game {
  var game Game
  DB.First(&game, id)

  if DB.NewRecord(game) {
    log.Fatalln("Could not find game with that ID")
  }

  return game
}

/* Calculates the Elo difference in rating between 
 * two players in a game 
 */
func diff(winnerRating int, loserRating int) int {
  outcome := 1.0
  expected := 1.0 / (1 + math.Pow(10, float64(loserRating - winnerRating) / 400.0))
  return int(K_FACTOR * (outcome - expected))
}

/* Create new NBA 2K game between two players*/
func (g Game) CreateNewGame() Game {
  var winner Player
  var loser Player
  winner = winner.GetById(g.WinnerId)
  loser = loser.GetById(g.LoserId)

  if g.WinnerPoints <= g.LoserPoints {
    log.Fatalln("Winner's score must be greater than loser's score")
  }

  rating_diff := diff(winner.Rating, loser.Rating)
  winner.Rating += rating_diff 
  loser.Rating -= rating_diff

  var game Game
  game.Date = int(time.Now().Unix())
  game.WinnerId = uint(winner.ID)
  game.LoserId = uint(loser.ID)
  game.WinnerRating = winner.Rating
  game.LoserRating = loser.Rating
  game.WinnerPoints = g.WinnerPoints
  game.LoserPoints = g.LoserPoints

  DB.Create(&game)
  DB.Save(&winner)
  DB.Save(&loser)

  return game
}
