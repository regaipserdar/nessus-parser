package utils

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite3 driver
)

// InsertData function to insert data into the database
func InsertData(db *sql.DB, tableName string, data map[string]interface{}) error {
	// Generate the SQL query for insertion dynamically based on the data map
	columns := ""
	placeholders := ""
	values := []interface{}{}
	i := 0
	for key, value := range data {
		if i > 0 {
			columns += ", "
			placeholders += ", "
		}
		columns += key
		placeholders += "?"
		values = append(values, value)
		i++
	}
	query := "INSERT INTO " + tableName + " (" + columns + ") VALUES (" + placeholders + ")"

	// Execute the SQL query
	_, err := db.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}
