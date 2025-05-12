package db

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "user-service/model"
)

var DB *gorm.DB

func Init() {
    var err error
    DB, err = gorm.Open(postgres.Open("host=postgres user=user password=password dbname=microservices_db port=5432 sslmode=disable"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    DB.AutoMigrate(&model.User{})
}
