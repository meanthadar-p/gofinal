package capi

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/meanthadar-p/gofinal/database/cdb"
	"github.com/meanthadar-p/gofinal/models"
)

func GetAllCustomerHandler(c *gin.Context){
	customers := []models.Customer{}
	customers = cdb.QueryAll()
	
	c.JSON(http.StatusOK, customers)
}
