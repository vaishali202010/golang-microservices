package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "product-service/db"
    "product-service/handlers"
)

func main() {
    db.Init()
    r := mux.NewRouter()
    r.HandleFunc("/product", handlers.Create).Methods("POST")
    r.HandleFunc("/product", handlers.GetAll).Methods("GET")
    r.HandleFunc("/product/{id}", handlers.GetOne).Methods("GET")
    r.HandleFunc("/product/{id}", handlers.Update).Methods("PUT")
    r.HandleFunc("/product/{id}", handlers.Delete).Methods("DELETE")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }
    log.Printf("product-service running on port %s", port)
    log.Fatal(http.ListenAndServe(":" + port, r))
}
