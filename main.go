package main

import (
  "../2KELO/routes"
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
