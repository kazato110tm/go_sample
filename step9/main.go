package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/tenntenn/sqlite"
)

func main() {
	db, err := sql.Open(sqlite.DriverName, "accountbook.db")
	if err != nil {
		log.Fatal(err)
	}

	ab := NewAccountBook(db)

	if err := ab.CreateTable(); err != nil {
		log.Fatal(err)
	}

	hs := NewHandlers(ab)

	http.HandleFunc("/", hs.ListHandler)

	fmt.Println("http://localhost:8080 で起動中...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
