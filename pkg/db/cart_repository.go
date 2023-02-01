package db

import "cart/pkg/models"

type CartRepository interface {
	Get(id int) (*models.Cart, error)
	GetAll(customerId string) ([]models.Cart, error)
	Create(c *models.Cart) error
	Update(c *models.Cart) error
	Delete(id int) error
}
