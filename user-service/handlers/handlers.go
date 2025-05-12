package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "user-service/db"
    "user-service/model"
)

func Create(w http.ResponseWriter, r *http.Request) {
    var m model.User
    json.NewDecoder(r.Body).Decode(&m)
    db.DB.Create(&m)
    json.NewEncoder(w).Encode(&m)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
    var items []model.User
    db.DB.Find(&items)
    json.NewEncoder(w).Encode(&items)
}

func GetOne(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var m model.User
    db.DB.First(&m, id)
    json.NewEncoder(w).Encode(&m)
}

func Update(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var m model.User
    db.DB.First(&m, id)
    json.NewDecoder(r.Body).Decode(&m)
    db.DB.Save(&m)
    json.NewEncoder(w).Encode(&m)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    db.DB.Delete(&model.User{}, id)
    w.WriteHeader(http.StatusNoContent)
}
