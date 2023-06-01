package services

import (
	"cart/internal/pkg/models"
	"log"
)

type MockCartRepository struct {
	Cart  *models.Cart
	Err   error
	Carts []models.Cart
}

func (r *MockCartRepository) Get(id int) (*models.Cart, error) {
	log.Println("MockCartRepository.Get called with id: ", id)
	return r.Cart, r.Err
}

func (r *MockCartRepository) GetAll(customerId string) ([]models.Cart, error) {
	log.Println("MockCartRepository.GetAll called with customerId: ", customerId)
	return r.Carts, r.Err
}

func (r *MockCartRepository) Create(c *models.Cart) error {
	log.Println("MockCartRepository.Create called with cart: ", c.ID)
	return r.Err
}

func (r *MockCartRepository) Update(c *models.Cart) error {
	log.Println("MockCartRepository.Update called with cart: ", c.ID)
	return r.Err
}

func (r *MockCartRepository) Delete(id int) error {
	log.Println("MockCartRepository.Delete called with id: ", id)
	return r.Err
}
