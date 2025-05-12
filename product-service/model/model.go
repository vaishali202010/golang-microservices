package model

type Product struct {
    ID    uint    `gorm:"primaryKey" json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}