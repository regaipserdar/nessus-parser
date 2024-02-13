package utils

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// CreateDB veritabanı bağlantısını oluşturur ve döndürür.
func CreateDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Printf("Error opening database: %v\n", err)
		return nil, err
	}

	// Veritabanı bağlantısını test et
	if err := db.Ping(); err != nil {
		log.Printf("Error connecting to database: %v\n", err)
		return nil, err
	}

	return db, nil
}
