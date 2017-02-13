package routes

import (
  "net/http"
  "encoding/json"
  "strconv"

  "github.com/gorilla/mux"

  "../models"
  "../errors"
)

const INITIAL_RATING = 1300

func PlayerRouter(r *mux.Router) {
  r.HandleFunc("/player", createPlayer).Methods("POST")
  r.HandleFunc("/player/{id}", getPlayerById).Methods("GET")
}

func createPlayer(rw http.ResponseWriter, req *http.Request) {
  decoder := json.NewDecoder(req.Body)

  var player models.Player
  err := decoder.Decode(&player)
  if err != nil {
    errors.ThrowError(rw, errors.InternalServerErr)
    return
  }

  defer req.Body.Close()

  player.Rating = INITIAL_RATING
  models.DB.Create(&player)
}

func getPlayerById(rw http.ResponseWriter, req *http.Request) {
  player_id, _ := strconv.Atoi(mux.Vars(req)["id"])
  var player models.Player

  models.GetByID(&player, uint(player_id))

  payload, _ := json.Marshal(player.ToJSON())
  
  rw.Write([]byte(payload))
}
