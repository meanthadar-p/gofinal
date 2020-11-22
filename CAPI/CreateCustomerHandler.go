package capi

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/meanthadar-p/gofinal/database/cdb"
	"github.com/meanthadar-p/gofinal/models"
)

func CreateCustomerHandler(c *gin.Context){
	customer := models.Customer{}
	if err := c.ShouldBindJSON(&customer); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	customer = cdb.Insert(customer.Name, customer.Email, customer.Status)
	
	c.JSON(http.StatusCreated, customer)
}
