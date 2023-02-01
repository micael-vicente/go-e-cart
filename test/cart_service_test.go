package services

import (
	"cart/pkg/models"
	"cart/pkg/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCart(t *testing.T) {
	expected := &models.Cart{CustomerId: "1"}
	repo := &MockCartRepository{Cart: expected}

	service := &services.CartService{Repo: repo}
	cart, err := service.CreateCart("1")
	assert.Nil(t, err)
	assert.Equal(t, expected.CustomerId, cart.CustomerId)
}

func TestUpdateCart(t *testing.T) {
	expected := &models.Cart{CustomerId: "1", CartItems: []models.CartItem{{Sku: "1", Quantity: 1}}}
	repo := &MockCartRepository{Cart: expected}

	service := &services.CartService{Repo: repo}
	cart, err := service.UpdateCart(1, []models.CartItem{{Sku: "1", Quantity: 1}})
	assert.Nil(t, err)
	assert.Equal(t, expected.CartItems, cart.CartItems)
}

func TestGetCart(t *testing.T) {
	expected := &models.Cart{ID: 1}
	repo := &MockCartRepository{Cart: expected}

	service := &services.CartService{Repo: repo}
	cart, err := service.GetCart(1)
	assert.Nil(t, err)
	assert.Equal(t, expected.ID, cart.ID)
}

func TestListByCustomer(t *testing.T) {
	expected := []models.Cart{{ID: 1}, {ID: 2}}
	repo := &MockCartRepository{Carts: expected}

	service := &services.CartService{Repo: repo}
	carts, err := service.ListByCustomer("1")
	assert.Nil(t, err)
	assert.Equal(t, expected, carts)
}
