package capi

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"

	"github.com/meanthadar-p/gofinal/database/cdb"
	"github.com/meanthadar-p/gofinal/models"
)

func UpdateCustomerHandler(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	customer := models.Customer{}
	if err := c.ShouldBindJSON(&customer); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer = cdb.Update(id, customer.Name, customer.Email, customer.Status)

	c.JSON(http.StatusOK, customer)
}
