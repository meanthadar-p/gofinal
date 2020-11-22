package cdb

import (
	"log"
	"github.com/meanthadar-p/gofinal/models"
)

func QueryAll()(customers []models.Customer){
	stmt, err := db.Prepare("SELECT id, name, email, status FROM customers")
	if err != nil {
		log.Fatal("can't prepare query all customers statment", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("can't query all customers", err)
	}
	
	for rows.Next() {
		var id int
		var name, email, status string

		if err := rows.Scan(&id, &name, &email, &status); err != nil {
			log.Fatal("can't Scan row into variable", err)
		}

		customer := models.Customer{
			ID: id,
			Name: name,
			Email: email,
			Status: status,
		}
		customers = append(customers, customer)
	}

	return customers
}
