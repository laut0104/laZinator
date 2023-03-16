package handler

import (
	"database/sql"

	_ "github.com/lib/pq"

	"fmt"
	"log"
)

func OpenDB() *sql.DB {
	connStr := "user=root dbname=lazinator password=password host=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
		return nil
	}
	fmt.Println("Open")

	return db
}
