package cdb

import (
	"log"
	"github.com/meanthadar-p/gofinal/models"
)

func Update(id int, name string, email string, status string) models.Customer{
	stmt, err := db.Prepare("UPDATE customers SET name=$2, email=$3, status=$4 WHERE id=$1;")

	if err != nil {
		log.Fatal("can't prepare statment update", err)
	}

	if _, err := stmt.Exec(id, name, email, status); err != nil {
		log.Fatal("error execute update ", err)
	}

	return models.Customer{
		ID: id,
		Name: name,
		Email: email,
		Status: status,
	}
}
