package cdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	DATABASE_URL = os.Getenv("DATABASE_URL")
) 

func init(){
	var err error
	db, err = sql.Open("postgres", DATABASE_URL)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	fmt.Println("Connected to database")
}

func Conn() *sql.DB{
	return db
}
