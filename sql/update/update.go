package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("update usuarios set nome = ?  where id = ?")
	stmt.Exec("Loren Ipsun", 1)
	stmt.Exec("Cheron St", 2)

	stmt2, err := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(3)
}
