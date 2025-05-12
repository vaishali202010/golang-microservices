package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "order-service/db"
    "order-service/handlers"
)

func main() {
    db.Init()
    r := mux.NewRouter()
    r.HandleFunc("/order", handlers.Create).Methods("POST")
    r.HandleFunc("/order", handlers.GetAll).Methods("GET")
    r.HandleFunc("/order/{id}", handlers.GetOne).Methods("GET")
    r.HandleFunc("/order/{id}", handlers.Update).Methods("PUT")
    r.HandleFunc("/order/{id}", handlers.Delete).Methods("DELETE")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }
    log.Printf("order-service running on port %s", port)
    log.Fatal(http.ListenAndServe(":" + port, r))
}
