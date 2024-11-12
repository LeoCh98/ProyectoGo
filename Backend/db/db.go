package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

// InitDB inicializa la conexión a la base de datos
func InitDB() {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/userdatabase")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
}

// CloseDB cierra la conexión a la base de datos
func CloseDB() {
	if db != nil {
		db.Close()
	}
}

// GetDB devuelve la conexión de la base de datos
func GetDB() *sql.DB {
	return db
}
