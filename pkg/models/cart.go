package models

type Cart struct {
	ID         int        `json:"id" gorm:"primary_key"`
	CustomerId string     `json:"customer_id"`
	CartItems  []CartItem `json:"items" gorm:"foreignkey:CartId" references:"id" cascade:"all,delete-orphan"`
}
