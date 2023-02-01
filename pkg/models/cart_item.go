package models

type CartItem struct {
	Id       int    `gorm:"primaryKey" json:"-"`
	CartId   int    `json:"-"`
	Sku      string `json:"sku"`
	Quantity int    `json:"quantity"`
}
