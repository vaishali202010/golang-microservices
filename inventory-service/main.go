package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "inventory-service/db"
    "inventory-service/handlers"
)

func main() {
    db.Init()
    r := mux.NewRouter()
    r.HandleFunc("/inventory", handlers.Create).Methods("POST")
    r.HandleFunc("/inventory", handlers.GetAll).Methods("GET")
    r.HandleFunc("/inventory/{id}", handlers.GetOne).Methods("GET")
    r.HandleFunc("/inventory/{id}", handlers.Update).Methods("PUT")
    r.HandleFunc("/inventory/{id}", handlers.Delete).Methods("DELETE")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }
    log.Printf("inventory-service running on port %s", port)
    log.Fatal(http.ListenAndServe(":" + port, r))
}
