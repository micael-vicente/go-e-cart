package db

import (
	"cart/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type CartRepositoryImpl struct {
	Db *gorm.DB
}

// Get returns a cart by id
func (r *CartRepositoryImpl) Get(id int) (*models.Cart, error) {
	var cart models.Cart
	log.Println("Getting cart with id: ", id)
	if err := r.Db.Preload(clause.Associations).First(&cart, id).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

// GetAll returns all carts for a customer
func (r *CartRepositoryImpl) GetAll(customerId string) ([]models.Cart, error) {
	var carts []models.Cart
	if err := r.Db.Where("customer_id = ?", customerId).Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

// Create saves a cart
func (r *CartRepositoryImpl) Create(cart *models.Cart) error {
	if err := r.Db.Create(cart).Error; err != nil {
		return err
	}
	return nil
}

// Update updates a cart
func (r *CartRepositoryImpl) Update(cart *models.Cart) error {
	err := r.Db.Model(&cart).Association("CartItems").Replace(cart.CartItems)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a cart
func (r *CartRepositoryImpl) Delete(id int) error {
	if err := r.Db.Where("id = ?", id).Delete(&models.Cart{}).Error; err != nil {
		return err
	}
	return nil
}
