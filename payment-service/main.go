package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "payment-service/db"
    "payment-service/handlers"
)

func main() {
    db.Init()
    r := mux.NewRouter()
    r.HandleFunc("/payment", handlers.Create).Methods("POST")
    r.HandleFunc("/payment", handlers.GetAll).Methods("GET")
    r.HandleFunc("/payment/{id}", handlers.GetOne).Methods("GET")
    r.HandleFunc("/payment/{id}", handlers.Update).Methods("PUT")
    r.HandleFunc("/payment/{id}", handlers.Delete).Methods("DELETE")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }
    log.Printf("payment-service running on port %s", port)
    log.Fatal(http.ListenAndServe(":" + port, r))
}
