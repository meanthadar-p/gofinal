package cdb

import (
	"fmt"
	"log"
)

func CreateTable(){
	createTb := `
	CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);
	`

	if _, err := db.Exec(createTb); err != nil {
		log.Fatal("can't create table", err)
	}

	fmt.Println("connected to database.")
}