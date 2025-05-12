package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "user-service/db"
    "user-service/handlers"
)

func main() {
    db.Init()
    r := mux.NewRouter()
    r.HandleFunc("/user", handlers.Create).Methods("POST")
    r.HandleFunc("/user", handlers.GetAll).Methods("GET")
    r.HandleFunc("/user/{id}", handlers.GetOne).Methods("GET")
    r.HandleFunc("/user/{id}", handlers.Update).Methods("PUT")
    r.HandleFunc("/user/{id}", handlers.Delete).Methods("DELETE")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }
    log.Printf("user-service running on port %s", port)
    log.Fatal(http.ListenAndServe(":" + port, r))
}
