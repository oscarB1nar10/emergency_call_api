// Go package
package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/lib/pq"
)

const (
    DB_USER = "postgres"
    DB_PASSWORD = "B1nar10.01"
    DB_NAME = "USER_TRACKING"
)

// DB set up
func setupDB() *sql.DB{
    dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)

    checkErr(err)

    return DB
}

// Go main function
func main() {
    fmt.Println("Hello World!")
}