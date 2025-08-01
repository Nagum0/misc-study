package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	dsn := "root:0000@(mysql-service:3306)/default_db"
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	var albums []Album
	err = db.Select(&albums, "SELECT * FROM Albums")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, album := range albums {
		fmt.Println(album)
	}
}
