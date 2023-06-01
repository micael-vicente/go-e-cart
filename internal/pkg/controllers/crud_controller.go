package controllers

import (
	"cart/internal/pkg/models"
	"cart/internal/pkg/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type CrudController struct {
	Service services.CartService
}

func (cc *CrudController) ListCarts(c *gin.Context) {
	carts, err := cc.Service.ListByCustomer(c.Query("customerId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, carts)
	}
}

func (cc *CrudController) GetCart(c *gin.Context) {
	cartId, _ := strconv.Atoi(c.Param("id"))
	log.Println("Getting cart with id: ", cartId)

	cart, err := cc.Service.GetCart(cartId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, cart)
	}
}

func (cc *CrudController) CreateCart(c *gin.Context) {
	log.Println("Creating cart for customer: ", c.Query("customerId"))
	cart, err := cc.Service.CreateCart(c.Query("customerId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, cart)
	}
}

func (cc *CrudController) DeleteCart(c *gin.Context) {
	cartId, _ := strconv.Atoi(c.Param("id"))
	log.Println("Deleting cart with id: ", cartId)
	err := cc.Service.DeleteCart(cartId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusNoContent, gin.H{"message": "Cart deleted"})
	}
}

func (cc *CrudController) UpdateCart(c *gin.Context) {
	cartId, _ := strconv.Atoi(c.Param("id"))
	log.Println("Updating cart with id: ", cartId)

	var items []models.CartItem
	if err := c.ShouldBindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	cart, err := cc.Service.UpdateCart(cartId, items)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, cart)
	}
}
