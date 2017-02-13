package config 

import (
  "os"
  "github.com/joho/godotenv"
)

var Env Environment

func init() {
  err := godotenv.Load()
  if err != nil {
    panic(err)
  }

  Env.DBUser = os.Getenv("DB_USER")
  Env.DBPassword = os.Getenv("DB_PASS")

}
