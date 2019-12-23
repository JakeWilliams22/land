package main

import (
    "database/sql"
)

var db *sql.DB
func init() {
  db = initDb()
}

func main() {
  defer db.Close()
  // testDb()
  testAPI()
  // startServer()
}