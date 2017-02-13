package middleware

import (
  "net/http"
  "fmt"
)

func ConfigResponseMiddleware(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
  rw.Header().Set("Content-Type", "application/json")
  rw.Header().Set("Access-Control-Allow-Origin", "*")
  rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
  rw.Header().Set("Access-Control-Allow-Headers",
    "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
  next(rw, req)
}

func HandleResponseMiddleware(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
  fmt.Println(req.Body)
}
