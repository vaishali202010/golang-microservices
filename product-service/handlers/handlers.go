package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "product-service/db"
    "product-service/model"
)

func Create(w http.ResponseWriter, r *http.Request) {
    var m model.Product
    json.NewDecoder(r.Body).Decode(&m)
    db.DB.Create(&m)
    json.NewEncoder(w).Encode(&m)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
    var items []model.Product
    db.DB.Find(&items)
    json.NewEncoder(w).Encode(&items)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var m model.Product
    db.DB.First(&m, id)
    json.NewEncoder(w).Encode(&m)
}

func Update(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var m model.Product
    db.DB.First(&m, id)
    json.NewDecoder(r.Body).Decode(&m)
    db.DB.Save(&m)
    json.NewEncoder(w).Encode(&m)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    db.DB.Delete(&model.Product{}, id)
    w.WriteHeader(http.StatusNoContent)
}
