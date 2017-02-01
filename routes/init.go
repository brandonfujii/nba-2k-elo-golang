package routes

import (
  "github.com/gorilla/mux"
  "github.com/urfave/negroni"
)

func Router() *negroni.Negroni {
  r  := mux.NewRouter()
  r.KeepContext = true
  r.StrictSlash(true)
  
  n := negroni.New()
  n.UseHandler(r)

  return n
}