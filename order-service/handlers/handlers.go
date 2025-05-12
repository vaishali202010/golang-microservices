package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "order-service/db"
    "order-service/model"
)

func Create(w http.ResponseWriter, r *http.Request) {
    var m model.Order
    json.NewDecoder(r.Body).Decode(&m)
    db.DB.Create(&m)
    json.NewEncoder(w).Encode(&m)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
    var items []model.Order
    db.DB.Find(&items)
    json.NewEncoder(w).Encode(&items)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var m model.Order
    db.DB.First(&m, id)
    json.NewEncoder(w).Encode(&m)
}

func Update(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var m model.Order
    db.DB.First(&m, id)
    json.NewDecoder(r.Body).Decode(&m)
    db.DB.Save(&m)
    json.NewEncoder(w).Encode(&m)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    db.DB.Delete(&model.Order{}, id)
    w.WriteHeader(http.StatusNoContent)
}
