package model

type Payment struct {
    ID      uint    `gorm:"primaryKey" json:"id"`
    OrderID uint    `json:"order_id"`
    Amount  float64 `json:"amount"`
}