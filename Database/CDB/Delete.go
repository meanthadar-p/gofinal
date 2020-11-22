package cdb

import (
	"fmt"
	"log"
)

func Delete(id int){
	stmt, err := db.Prepare("DELETE FROM customers WHERE id = $1")
	if err != nil {
		log.Fatal("can't prepare delete statement", err)
	}

	if _, err := stmt.Exec(fmt.Sprint(id)); err != nil {
		log.Fatal("can't execute delete statment", err)
	}

	fmt.Println("delete success")
}
