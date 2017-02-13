package routes

import (
  "net/http"
  "github.com/gorilla/mux"
  "github.com/urfave/negroni"
)

func Router() *negroni.Negroni {
  n := negroni.New()

  n.Use(negroni.HandlerFunc(initResponseMiddleware))

  r := mux.NewRouter()
  r.KeepContext = true
  r.StrictSlash(true)

  PlayerRouter(r)
  GameRouter(r)
  n.UseHandler(r)

  return n
}

/* Middleware */ 
func initResponseMiddleware(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
  rw.Header().Set("Content-Type", "application/json")
  rw.Header().Set("Access-Control-Allow-Origin", "*")
  rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
  rw.Header().Set("Access-Control-Allow-Headers",
    "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
 next(rw, req)
}
