package routes

import (
  "github.com/gorilla/mux"
  "github.com/urfave/negroni"
  "../middleware"
)

func Router() *negroni.Negroni {
  n := negroni.New()

  n.Use(negroni.NewLogger())
  n.Use(negroni.HandlerFunc(middleware.ConfigResponseMiddleware))

  r := mux.NewRouter()
  r.KeepContext = true
  r.StrictSlash(true)

  PlayerRouter(r)
  GameRouter(r)
  n.UseHandler(r)

  n.Use(negroni.HandlerFunc(middleware.HandleResponseMiddleware))

  return n
}