package model

type Inventory struct {
    ID              uint `gorm:"primaryKey" json:"id"`
    ProductID       uint `json:"product_id"`
    AvailableAmount int  `json:"available_amount"`
}