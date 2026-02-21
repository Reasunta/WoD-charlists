package main

import (
    "fmt"
    "net/http"
    "wod-backend/db"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        result, err := db.Now()
        if err != nil {
            fmt.Fprintf(w, "Unable to read a database: %v\n", err)
            return
        }

        fmt.Fprintf(w, "%v\n", result)
    })

    err := http.ListenAndServe(":3030", nil)
    if err != nil {
        fmt.Println("Error starting the server:", err)
    }
}