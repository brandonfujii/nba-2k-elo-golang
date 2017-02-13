# NBA 2K Player Rating API
A Go API that rates NBA 2K players based on head-to-head games

It uses Gorilla Mux as an HTTP router, GORM as an object relational mapper, and an variation of the Elo algorithm as a means of rating players

### Getting Started
  Create local MySQL database instance called `2K`
    `create database 2K`

  Clone and create a `.env` file in the root project directory, containing `DB_USER` (MySQL database user) and `DB_PASS` (corresponding password).

  Run `go run bin/create_db.go`, which should generate tables in your newly created database

  Install all necessary go dependencies and `go run main.go`

