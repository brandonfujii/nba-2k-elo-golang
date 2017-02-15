package main

import (
  "../nba-2k-elo-golang/routes"
  _ "net/http/pprof"
  "net/http"
  "log"
)

func main() {
  go func() {
    log.Println(http.ListenAndServe(":6060", nil))
  }()
  n := routes.Router()
  n.Run(":8080")
}
