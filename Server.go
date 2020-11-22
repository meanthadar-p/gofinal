package main

import (
	"github.com/gin-gonic/gin"

	CAPI "github.com/meanthadar-p/gofinal/capi"
)

func setupRouter() *gin.Engine{
	r := gin.Default()
	r.POST("/customers", CAPI.CreateCustomerHandler)
	r.PUT("/customers/:id", CAPI.UpdateCustomerHandler)
	r.GET("/customers", CAPI.GetAllCustomerHandler)
	r.GET("/customers/:id", CAPI.GetCustomerByIdHandler)
	r.DELETE("customers/:id", CAPI.DeleteCustomerHandler)

	return r
}

func main(){
	r := setupRouter()
	r.Run(":2009")
}