package capi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/meanthadar-p/gofinal/database/cdb"
	"github.com/meanthadar-p/gofinal/models"
)

func GetCustomerByIdHandler(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	customer := models.Customer{}
	customer = cdb.QueryById(id)

	c.JSON(http.StatusOK, customer)
}
