package errors

import (
  "net/http"
  "encoding/json"
)

type HttpError struct {
  Status  int    `json: "status"`
  Title   string `json: "title"`
  Message string `json: "message"`
}

func (err *HttpError) ToJSON() interface{} {
  data := make(map[string]interface{})
  data["status"] = err.Status
  data["title"] = err.Title
  data["message"] = err.Message
  return data
}

func ThrowError(rw http.ResponseWriter, err *HttpError) {
  rw.Header().Set("Content-Type", "application/json")
  rw.WriteHeader(err.Status)
  json.NewEncoder(rw).Encode(err.ToJSON())
}

var (
  BadRequestErr           = &HttpError{400, "Bad Request", "Request was faulty"}
  UnsupportedMediaTypeErr = &HttpError{415, "Unsupported Media Type", "Content-Type header must be set to: 'application/json'"}
  InternalServerErr       = &HttpError{500, "Internal Server Error", "Sorry. An unknown error has occurred"}
)
