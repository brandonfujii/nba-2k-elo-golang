package routes

import (
  "fmt"
  "strconv"
  "net/http"
  "encoding/json"

  "github.com/gorilla/mux"
  "../models"
)

func GameRouter(r *mux.Router) {
  r.HandleFunc("/games", createGame).Methods("POST")
  r.HandleFunc("/games/{id}", getGameById).Methods("GET")
}

func createGame(rw http.ResponseWriter, req *http.Request) {
  decoder := json.NewDecoder(req.Body)

  var game models.Game
  err := decoder.Decode(&game)
  if err != nil {
    panic(err)
    return
  }
  defer req.Body.Close()

  fmt.Println(game)

  game = game.CreateNewGame()
}

func getGameById(rw http.ResponseWriter, req *http.Request) {
  game_id, _ := strconv.Atoi(mux.Vars(req)["id"])
  var game models.Game

  game = game.GetById(uint(game_id))
  j, _ := json.Marshal(game.ToJSON())

  rw.Write(j)
}
