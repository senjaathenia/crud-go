package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB()  {
	//Koneksi ke Database SQL
	db, err := sql.Open("mysql", "root:shiroyasha@tcp(localhost:3306)/products?parseTime=true")
	if err != nil {
		panic(err)
	}
	log.Println("Database Connected")
	DB = db
}