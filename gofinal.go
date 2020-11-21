package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	//"fmt"
)

type Customer struct {
	id		int
	name	string
	email	string
	status	string
}

var customers = map[string]*Customer{}

func createCustomerHandler(c *gin.Context){

}

func updateCustomerHandler(c *gin.Context){

}

func getAllCustomerHandler(c *gin.Context){
	lists := []*Customer{}

	for _, list := range customers{
		lists = append(lists, list)
	}

	c.JSON(http.StatusOK, lists)
}

func getCustomerByIdHandler(c *gin.Context){

}

func deleteCustomerHandler(c *gin.Context){

}

func setupRouter() *gin.Engine{
	r := gin.Default()
	r.POST("/customers", createCustomerHandler)
	r.PUT("/customers/:id", updateCustomerHandler)
	r.GET("/customers", getAllCustomerHandler)
	r.GET("/customers/:id", getCustomerByIdHandler)
	r.DELETE("customers/:id", deleteCustomerHandler)

	return r
}

func main(){
	r := setupRouter()
	r.Run(":2009")
}