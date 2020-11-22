package capi

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"

	"github.com/meanthadar-p/gofinal/database/cdb"
	"github.com/meanthadar-p/gofinal/models"
)

func DeleteCustomerHandler(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	cdb.Delete(id)
	
	c.JSON(http.StatusOK, models.Message{Message:"customer deleted"})
}
