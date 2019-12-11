package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:postgres@localhost/hello?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT * FROM public.items`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var ntinid int
		var serial string
		var status int

		err = rows.Scan(&id, &ntinid, &serial, &status)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%d - %d - %s - %d\n", id, ntinid, serial, status)
	}
}
