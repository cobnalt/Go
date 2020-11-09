package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type product struct {
	id   int
	name string
}

func main() {
	connStr := "user=postgres password=docker dbname=test_base sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//_, err = db.Exec("insert into test_tbl (name) values ('Паша')")
	//if err != nil {
	//	panic(err)
	//}

	rows, err := db.Query("select * from test_tbl")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []product{}

	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.id, &p.name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	for _, p := range products {
		fmt.Println(p.id, p.name)
	}
}
