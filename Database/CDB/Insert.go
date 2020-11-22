package cdb

import (
	"log"
	"github.com/meanthadar-p/gofinal/models"
)

func Insert(name string, email string, status string) models.Customer {
	row := db.QueryRow("INSERT INTO customers (name, email, status) values ($1, $2, $3) RETURNING id", name, email, status)
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Fatal("can't prepare statment insert", err)
	}

	return models.Customer{
		ID: id,
		Name: name,
		Email: email,
		Status: status,
	}
}
