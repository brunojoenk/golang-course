package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:admin123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, nome from usuarios where id > ?", 0)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}
}
