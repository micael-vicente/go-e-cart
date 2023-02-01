package services

import (
	"cart/pkg/db"
	"cart/pkg/models"
)

type CartService struct {
	Repo db.CartRepository
}

// CreateCart creates a new cart
func (cs *CartService) CreateCart(customerId string) (*models.Cart, error) {
	var cart = models.Cart{CustomerId: customerId}
	return &cart, cs.Repo.Create(&cart)
}

// UpdateCart updates a cart
func (cs *CartService) UpdateCart(id int, items []models.CartItem) (*models.Cart, error) {
	if cart, err := cs.Repo.Get(id); err != nil {
		return nil, err
	} else {
		cart.CartItems = []models.CartItem{}
		cart.CartItems = append(cart.CartItems, items...)
		return cart, cs.Repo.Update(cart)
	}
}

// GetCart returns a cart by id
func (cs *CartService) GetCart(id int) (*models.Cart, error) {
	return cs.Repo.Get(id)
}

// ListByCustomer returns all carts for a customer
func (cs *CartService) ListByCustomer(customerId string) ([]models.Cart, error) {
	return cs.Repo.GetAll(customerId)
}

// DeleteCart deletes a cart
func (cs *CartService) DeleteCart(id int) error {
	return cs.Repo.Delete(id)
}
