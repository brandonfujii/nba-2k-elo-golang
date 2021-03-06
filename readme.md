# NBA 2K Player Rating API
A Go API that rates NBA 2K players based on head-to-head games

It uses Gorilla Mux as an HTTP request multiplexer, GORM as an object relational mapper, and a variation of the Elo algorithm as a means of rating players

### Getting Started
  Create local MySQL database instance
    `create database my_2k_db`

  Clone and create an `.env` file in the root project directory, containing `DB_NAME` (name of your database), `DB_USER` (a valid MySQL database user) and `DB_PASS` (corresponding password)

  Run `go run bin/create_db.go`, which should generate tables in your newly created database

  Install all necessary go dependencies and `go run main.go`

