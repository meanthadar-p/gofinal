package cdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func DropTable() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	dropTb := "DROP TABLE IF EXISTS customers;"
	if _, err := db.Exec(dropTb); err != nil {
		log.Fatal("can't drop table", dropTb)
	}

	fmt.Println("drop table success")
}
