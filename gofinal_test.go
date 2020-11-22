package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRouter(t *testing.T){
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/customers", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}

func TestCreateCustomer(t *testing.T){
	body := bytes.NewBufferString(`{
		"name": "Meanthadar",
		"email": "Meanthadar.p@gmail.com",
		"status": "active"
	}`)

	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/customers", body)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "{\"id\":\"1\",\"name\":\"Meanthadar\",\"email\":\"Meanthadar.p@gmail.com\",\"status\":\"active\"}", w.Body.String())
}