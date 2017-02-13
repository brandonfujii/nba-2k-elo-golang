package routes

import (
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "../models"
  "fmt"
)

func GameRouter(r *mux.Router) {
  r.HandleFunc("/games", createGame).Methods("POST")
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
