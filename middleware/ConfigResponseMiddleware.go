package middleware

import (
  "net/http"
)

func ConfigResponseMiddleware(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
  rw.Header().Set("Content-Type", "application/json")
  rw.Header().Set("Access-Control-Allow-Origin", "*")
  rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
  rw.Header().Set("Access-Control-Allow-Headers",
    "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
 next(rw, req)
}
