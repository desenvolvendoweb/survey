package main

import (
    "database/sql"
    "fmt"
    //"log"
    "os"

    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

func Connect() (db *sql.DB) {

    if godotenv.Load() != nil {
        fmt.Println("Error loading .env file")
    }

    dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_SERVICE"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASS"),
        os.Getenv("DB_NAME"))
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    return
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
