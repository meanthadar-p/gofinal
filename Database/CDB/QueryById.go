package cdb

import (
	"log"
	"github.com/meanthadar-p/gofinal/models"
)

func QueryById(id int) models.Customer {
	stmt, err := db.Prepare("SELECT id, name, email, status FROM customers where id=$1")
	if err != nil {
		log.Fatal("can'tprepare query one row statment", err)
	}

	row := stmt.QueryRow(id)
	var name, email, status string

	err = row.Scan(&id, &name, &email, &status)
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}

	return models.Customer{
		ID: id,
		Name: name,
		Email: email,
		Status: status,
	}
}